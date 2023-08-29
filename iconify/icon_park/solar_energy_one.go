package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SolarEnergyOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path fill="#2F88FF" stroke-linejoin="round" d="M10 30H16C19.3137 30 22 32.6863 22 36C22 39.3137 19.3137 42 16 42H10V30Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M10 28V44"/><path stroke-linecap="round" stroke-linejoin="round" d="M4 32H10"/><path stroke-linecap="round" stroke-linejoin="round" d="M4 38H10"/><path stroke-linecap="round" stroke-linejoin="round" d="M22 36H30C31.1046 36 32 35.1046 32 34V23"/><path stroke-linecap="round" stroke-linejoin="round" d="M32 4V11"/><path stroke-linecap="round" stroke-linejoin="round" d="M22.8789 7.87891L27.1215 12.1215"/><path stroke-linecap="round" stroke-linejoin="round" d="M22.8789 25.1211L27.1215 20.8785"/><path stroke-linecap="round" stroke-linejoin="round" d="M41.1211 7.87891L36.8785 12.1215"/><path stroke-linecap="round" stroke-linejoin="round" d="M41.1211 25.1211L36.8785 20.8785"/><circle cx="32" cy="17" r="6" fill="#2F88FF"/><path stroke-linecap="round" stroke-linejoin="round" d="M20 17H26"/><path stroke-linecap="round" stroke-linejoin="round" d="M38 17H44"/></g>`),
		g.Group(children),
	)
}