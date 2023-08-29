package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nekogotchi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.597 4.527c1.026 3.63.63 5.983-.662 8.446c-.227.422-5.916-2.428-5.63-2.78c2.765-1.778 3.634-2.803 6.292-5.666ZM10.406 4.5c-1.03 3.63-.635 5.982.661 8.445c.227.423 5.916-2.431 5.627-2.78c-2.76-1.777-3.634-2.803-6.288-5.665Zm18.781 16.424a6.8 6.8 0 1 1-6.942-6.796a6.967 6.967 0 0 1 6.942 6.796Zm1.903 16.59c0 3.29-3.76 5.963-8.418 5.986s-8.468-2.612-8.535-5.9c-.065-3.29 3.64-6 8.296-6.069s8.524 2.695 8.657 5.982Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.185 25.297c-2.278 2.099-2.267 4.837 0 7.62m15.808-17.513a45.179 45.179 0 0 0-10.556-1.278"/>`),
		g.Group(children),
	)
}