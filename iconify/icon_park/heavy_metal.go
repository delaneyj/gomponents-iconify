package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeavyMetal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linejoin="round" d="M22.2187 7.37842L33.8868 16.6216L36.1778 23.9998L25.2625 26.0852L12.5889 14.5312L15.281 9.00148L22.2187 7.37842Z" clip-rule="evenodd"/><path stroke-linecap="round" d="M15.2808 9.00146L26.4862 18.6016"/><path stroke-linecap="round" stroke-linejoin="round" d="M25.2627 26.0852L26.4866 18.1324L33.887 16.6216"/><path stroke-linecap="round" stroke-linejoin="round" d="M24.993 31.0228L27.2792 38.1032L15.601 40.3785L4 29.014L6.57052 22.6253L12.0285 21.3535"/><path stroke-linecap="round" d="M6.57031 22.625L17.2837 32.7577"/><path stroke-linecap="round" stroke-linejoin="round" d="M15.6011 39.8648L17.4865 32.1324L24.9908 31.0435"/><path stroke-linejoin="round" d="M34.8868 29.6077L34 36.7999L43.236 34.9985L41.2813 28.1875L34.8868 29.6077Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}