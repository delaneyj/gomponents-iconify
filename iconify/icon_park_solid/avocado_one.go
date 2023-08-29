package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AvocadoOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAvocadoOne0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path stroke="#fff" d="M7 32c3 5 8 8 13 9s10.808.144 15-3c6-4.5 8-12 8-17s-2-10.5-3-12s-3-2-3-2"/><path fill="#fff" stroke="#fff" d="M21 13c-4.355 1.604-9 0-13 3c-4.696 3.523-6 9-3 14c2.352 3.92 6 6 12 6s9.764-1.765 14-6c6-6 8-15.5 8-19s-2.5-5.5-6-5s-6.427 4.947-12 7Z"/><path fill="#000" stroke="#000" d="M21.945 20.117c-1.223.48-2.527 0-3.65.898c-1.32 1.054-1.686 2.693-.843 4.19c.66 1.172 1.685 1.795 3.37 1.795s2.742-.528 3.931-1.796C26.438 23.41 27 20.566 27 19.52c0-1.047-.702-1.646-1.685-1.496c-.983.15-1.805 1.48-3.37 2.094Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAvocadoOne0)"/>`),
		g.Group(children),
	)
}