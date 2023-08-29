package iconify

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"text/template"

	"github.com/delaneyj/toolbelt"
	"github.com/divan/num2words"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/goccy/go-json"
	"github.com/valyala/bytebufferpool"
	"golang.org/x/exp/slices"
	"golang.org/x/sync/errgroup"

	_ "embed"
)

//go:embed generateIcon.gotmpl
var generateIconTmpl string

//go:embed generateIconNamespace.gotmpl
var generateIconNamespaceTmpl string

func GenerateIconify(ctx context.Context, gentmpDir, output string) error {
	generateIconTmpl := template.Must(template.New("generateIcon").Parse(generateIconTmpl))
	generateIconNamespaceTmpl := template.Must(template.New("generateIconNamespace").Parse(generateIconNamespaceTmpl))

	tmpParentDir := filepath.Join(gentmpDir, "iconify")
	tmpIconSetsDir := filepath.Join(tmpParentDir, "json")
	if _, err := os.Stat(filepath.Join(tmpParentDir, "collections.json")); os.IsNotExist(err) {
		if err := os.MkdirAll(tmpIconSetsDir, 0755); err != nil {
			return fmt.Errorf("could not create iconify directory: %w", err)
		}

		iconCollections := map[string]iconBasicCollectionInfo{}
		if err := updateIconifyCache(ctx, tmpParentDir, "collections.json", &iconCollections); err != nil {
			return fmt.Errorf("could not get iconify collections: %w", err)
		}
		log.Print("got icon collections")

		eg, ctx := errgroup.WithContext(ctx)

		for collectionName := range iconCollections {
			details := iconCollectionDetailsInfo{}
			collectionName := collectionName

			eg.Go(func() error {
				path := fmt.Sprintf("json/%s.json", collectionName)
				if err := updateIconifyCache(ctx, tmpParentDir, path, &details); err != nil {
					log.Printf("!!!! could not get iconify collection %s: %v", collectionName, err)
				} else {
					log.Printf("got icon collection %s", collectionName)
				}
				return nil
			})
		}

		if err := eg.Wait(); err != nil {
			return fmt.Errorf("could not get iconify collections: %w", err)
		}
	}

	iconifyPath := filepath.Join(output, "iconify")
	if err := os.RemoveAll(iconifyPath); err != nil {
		return fmt.Errorf("could not remove iconify directory: %w", err)
	}
	if err := os.MkdirAll(iconifyPath, 0755); err != nil {
		return fmt.Errorf("could not create iconify directory: %w", err)
	}

	numberWordGroupsRegex := regexp.MustCompile(`(?P<number>\d*)*(?P<word>\D*)`)

	// walk the details and generate the iconify.go file
	if err := filepath.WalkDir(tmpIconSetsDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() || filepath.Ext(path) != ".json" {
			return nil
		}

		b, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("could not read file %s: %w", path, err)
		}

		details := iconCollectionDetailsInfo{}
		if err := json.Unmarshal(b, &details); err != nil {
			return fmt.Errorf("could not unmarshal JSON: %w", err)
		}

		// some icons have a prefix that is not a valid Go identifier.  Specifically, the
		// use of numbers, dashes, and underscores.  We need to convert these to valid
		// Go identifiers.

		namesUsed := sets.NewString()
		namedIcons := []NamedIcon{}
		for iconName, icon := range details.Icons {
			matches := numberWordGroupsRegex.FindAllStringSubmatch(iconName, -1)
			revisedName := iconName

			if len(matches) > 0 {
				parts := []string{}

				for _, match := range matches {
					numbersStr := match[numberWordGroupsRegex.SubexpIndex("number")]
					if numbersStr != "" {
						number, err := strconv.Atoi(numbersStr)
						if err != nil {
							return fmt.Errorf("could not convert %s to number: %w", numbersStr, err)
						}
						numbersStr = num2words.Convert(number) + "_"
					}

					word := match[numberWordGroupsRegex.SubexpIndex("word")]
					parts = append(parts, numbersStr, word)
				}

				revisedName = toolbelt.Pascal(strings.Join(parts, " "))
			}

			if namesUsed.Has(revisedName) {
				continue
			}

			namedIcons = append(namedIcons, NamedIcon{
				OriginalName: iconName,
				Name:         toolbelt.Cased(revisedName),
				Icon:         icon,
			})
			namesUsed.Insert(revisedName)
		}

		slices.SortFunc(namedIcons, func(a, b NamedIcon) bool {
			return a.Name.Original < b.Name.Original
		})

		packageName := toolbelt.Lower(toolbelt.Snake(details.Prefix))
		if goKeywords.Has(packageName) {
			packageName = fmt.Sprintf("%s_icons", packageName)
		}

		iconPkg := IconPackage{
			Name:       toolbelt.Cased(packageName),
			Icons:      namedIcons,
			Width:      unknownToDimension(details.Info.Height),
			Height:     unknownToDimension(details.Info.Height),
			ViewWidth:  unknownToDimension(details.Width),
			ViewHeight: unknownToDimension(details.Height),
			Version:    details.Info.Version,
		}
		if iconPkg.Width == 0 && iconPkg.Height > 0 {
			iconPkg.Width = iconPkg.Height
		} else if iconPkg.Height == 0 && iconPkg.Width > 0 {
			iconPkg.Height = iconPkg.Width
		} else if iconPkg.Width == 0 && iconPkg.Height == 0 {
			iconPkg.Width = 24
			iconPkg.Height = 24
		}

		packagePath := filepath.Join(iconifyPath, packageName)
		if err := os.MkdirAll(packagePath, 0755); err != nil {
			return fmt.Errorf("could not create iconify directory: %w", err)
		}
		namespaceFullPath := filepath.Join(packagePath, fmt.Sprintf("%s_pkg_info.go", packageName))

		namespaceFile, err := os.Create(namespaceFullPath)
		if err != nil {
			return fmt.Errorf("could not create file %s: %w", namespaceFullPath, err)
		}
		defer namespaceFile.Close()

		if err := generateIconNamespaceTmpl.Execute(namespaceFile, iconPkg); err != nil {
			return fmt.Errorf("could not execute template: %w", err)
		}

		type IconCtx struct {
			Name                  toolbelt.CasedString
			PackageName           toolbelt.CasedString
			SvgBody               string
			ViewWidth, ViewHeight int
		}

		wg := &sync.WaitGroup{}
		wg.Add(len(iconPkg.Icons))
		errCh := make(chan error, len(iconPkg.Icons))

		for _, icon := range iconPkg.Icons {
			func(icon NamedIcon) {
				defer wg.Done()

				iconFullPath := filepath.Join(packagePath, fmt.Sprintf("%s.go", icon.Name.Snake))
				iconFile, err := os.Create(iconFullPath)
				if err != nil {
					errCh <- fmt.Errorf("could not create file %s: %w", iconFullPath, err)
					return
				}
				defer iconFile.Close()

				iconCtx := IconCtx{
					PackageName: iconPkg.Name,
					Name:        icon.Name,
					SvgBody:     icon.Icon.SvgBody,
					ViewWidth:   iconPkg.ViewWidth,
					ViewHeight:  iconPkg.ViewHeight,
				}

				buf := bytebufferpool.Get()
				defer bytebufferpool.Put(buf)

				err = generateIconTmpl.Execute(buf, iconCtx)
				if err != nil {
					err = fmt.Errorf("could not execute template: %w", err)
					errCh <- err
					return
				}

				if _, err := iconFile.Write(buf.Bytes()); err != nil {
					err = fmt.Errorf("could not write file %s: %w", iconFullPath, err)
					errCh <- err
					return
				}

			}(icon)
		}
		wg.Wait()

		if len(errCh) > 0 {
			errs := []error{}
			for err := range errCh {
				errs = append(errs, err)
			}
			return errors.Join(errs...)
		}

		log.Printf("wrote %s", namespaceFullPath)

		return nil
	}); err != nil {
		return fmt.Errorf("could not walk iconify directory: %w", err)
	}

	return nil
}

var goKeywords = sets.New(
	"break", "case", "chan", "const", "continue", "default", "defer",
	"else", "fallthrough", "for", "func", "go", "goto", "if", "import",
	"interface", "map", "package", "range", "return", "select", "struct",
	"switch", "type", "var",
)

func updateIconifyCache(ctx context.Context, gentmpDir, iconifyJSONPath string, v interface{}) error {
	fullURL := fmt.Sprintf("https://raw.githubusercontent.com/iconify/icon-sets/master/%s", iconifyJSONPath)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return fmt.Errorf("could not create request: %w", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not get iconify %s: %w", fullURL, err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("could not get iconify %s: %v", fullURL, res.Status)
	}

	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)

	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		return fmt.Errorf("could not read body: %w", err)
	}

	b := buf.Bytes()

	if err := json.Unmarshal(b, v); err != nil {
		return fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	jsonPath := filepath.Join(gentmpDir, iconifyJSONPath)
	subDirPath := filepath.Dir(jsonPath)
	if err := os.MkdirAll(subDirPath, 0755); err != nil {
		return fmt.Errorf("could not create directory %s: %w", subDirPath, err)
	}

	bytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("could not marshal JSON: %w", err)
	}
	if err := os.WriteFile(jsonPath, bytes, 0644); err != nil {
		return fmt.Errorf("could not write file %s: %w", jsonPath, err)
	}

	return nil
}

type Author struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type License struct {
	Title string `json:"title,omitempty"`
	SPDX  string `json:"spdx,omitempty"`
	URL   string `json:"url,omitempty"`
}

type iconBasicCollectionInfo struct {
	Name     string  `json:"name,omitempty"`
	Total    int     `json:"total,omitempty"`
	Version  string  `json:"version,omitempty"`
	Author   Author  `json:"author,omitempty"`
	License  License `json:"license,omitempty"`
	Category string  `json:"category,omitempty"`
}

type iconCollectionDetailsInfo struct {
	Prefix string `json:"prefix,omitempty"`
	Info   struct {
		Name     string   `json:"name,omitempty"`
		Total    int      `json:"total,omitempty"`
		Version  string   `json:"version,omitempty"`
		Author   Author   `json:"author,omitempty"`
		License  License  `json:"license,omitempty"`
		Samples  []string `json:"samples,omitempty"`
		Height   int      `json:"height,omitempty"`
		Category string   `json:"category,omitempty"`
		Palette  bool     `json:"palette,omitempty"`
	} `json:"info,omitempty"`
	LastModified int64             `json:"lastModified,omitempty"`
	Icons        map[string]Icon   `json:"icons,omitempty"`
	Suffixes     map[string]string `json:"suffixes,omitempty"`
	Width        interface{}       `json:"width,omitempty"`
	Height       interface{}       `json:"height,omitempty"`
}

func unknownToDimension(x interface{}) int {
	switch v := x.(type) {
	case int:
		return v
	case float64:
		return int(v)
	case []int:
		return v[0]
	default:
		return 0
	}
}

type IconPackage struct {
	Name                  toolbelt.CasedString
	Icons                 []NamedIcon
	Width, Height         int
	ViewWidth, ViewHeight int
	Version               string
}

type NamedIcon struct {
	OriginalName string
	Name         toolbelt.CasedString
	Icon         Icon
}

type Icon struct {
	SvgBody string      `json:"body,omitempty"`
	Width   interface{} `json:"width,omitempty"`
	Height  interface{} `json:"height,omitempty"`
}
