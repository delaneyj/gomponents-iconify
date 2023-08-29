package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LipTattoo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSLipTattoo0"><g fill="none" stroke-width="4"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M25 19.89c-1.5-2.52-5.5-2.52-7 0c-1.115 1.874-5 4.61-6 5.053C13.667 27.295 19.5 32 25 32c6.5 0 11.167-4.704 13-7.057c-1-.443-3.164-2.123-5.5-5.053c-2-2.508-6-2.508-7.5 0Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M12 25c3.79.755 14.296 1.811 26 0"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M44 24c0 11.046-8.954 20-20 20S4 35.046 4 24S12.954 4 24 4"/><path stroke="#fff" d="m39 8.472l.343 1.056h1.11l-.898.652l.343 1.056l-.898-.652l-.898.652l.343-1.056l-.898-.652h1.11L39 8.472Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M35 22.656c1.294 1.248 2.374 2.01 3 2.287A20.974 20.974 0 0 1 33.763 29m-17.665-7c-1.543 1.374-3.442 2.652-4.098 2.943c.941 1.329 3.211 3.407 6 4.965"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSLipTattoo0)"/>`),
		g.Group(children),
	)
}