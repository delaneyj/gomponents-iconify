package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Palette(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="18.314" height="39" x="14.843" y="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3"/><path fill="none" stroke="currentColor" d="M14.843 33.4h18.314M14.843 22.235s3.933-.232 5.292 2.271s2.605 3.387 4.466.291s3.5-6.83 8.556-.514m-.001-11.601s-1.669-2.183-3.418-2.088s-4.099 4.972-5.418 4.852s-2.285-6.481-4.118-7.03s-5.36 2.35-5.36 2.35"/><ellipse cx="24" cy="38.01" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.965" ry="1.957"/>`),
		g.Group(children),
	)
}