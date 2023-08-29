package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdobeCharacteranimator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 0h512v512H0V0zm422.344 240.566v123.662h-41.858V244.79c1.322-42.628-37.786-42.766-58.76-24.197v143.634h-41.862V108.07h41.862v78.73c30.756-23.923 102.116-28.254 100.618 53.767zM246.836 110.475v40.266c-59.386-24.703-116.64 2.886-116.75 84.446c1.644 69.545 38.569 109.902 114.446 88.16l.002 37.578C154.1 387.402 84.891 335.774 86.304 236.342c.633-98.056 69.098-152.01 160.532-125.867z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}