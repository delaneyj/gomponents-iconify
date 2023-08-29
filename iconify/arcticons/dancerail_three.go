package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DancerailThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.6 27.7h3.7v3.7h-3.7zm-5.55 0h3.7v3.7h-3.7zm-5.55 0h3.7v3.7H5.5zm0-5.55h3.7v3.7H5.5zM22.15 5.5l3.7 3.7h-3.7V5.5Zm5.55 11.1h3.7v3.7h-3.7zm5.55 0h3.7v3.7h-3.7zm-11.1 5.55h3.7v3.7h-3.7zm0 5.55h3.7v3.7h-3.7zm5.55 0h3.7v3.7h-3.7zm5.55 0h3.7v3.7h-3.7zm5.55 11.1h3.7v3.7h-3.7zm-16.65-5.55h3.7v3.7h-3.7zm0 5.55h3.7v3.7h-3.7zm16.65-5.55h3.7v3.7h-3.7zM5.5 16.6h3.7v3.7H5.5zm16.65-5.55h3.7v3.7h-3.7zm0 5.55h3.7v3.7h-3.7zM5.5 5.5h3.7v3.7H5.5zm0 5.55h3.7v3.7H5.5zm33.3 5.55l3.7 3.7h-3.7v-3.7Zm0 5.55h3.7v3.7h-3.7zM11.05 5.5h3.7v3.7h-3.7zm5.55 0h3.7v3.7h-3.7z"/>`),
		g.Group(children),
	)
}