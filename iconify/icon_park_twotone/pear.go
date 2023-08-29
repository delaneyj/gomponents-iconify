package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPear0"><g fill="none"><path fill="#555" stroke="#fff" stroke-width="4" d="M16.64 12.746a8.081 8.081 0 0 1 14.72 0l3.061 6.754l3.686 6.322a9.88 9.88 0 0 1-1.652 12.063l-1.19 1.156a8.013 8.013 0 0 1-10.634.472a1.002 1.002 0 0 0-1.262 0a8.013 8.013 0 0 1-10.634-.472l-1.19-1.156a9.88 9.88 0 0 1-1.652-12.063L13.58 19.5l3.06-6.754Z"/><path stroke="#fff" stroke-linecap="round" stroke-width="4" d="M21.5 4c.667.5 2.1 1.9 2.5 3.5c.4 1.6.167 3.333 0 4"/><circle cx="16" cy="28" r="2" fill="#fff"/><circle cx="18" cy="34" r="2" fill="#fff"/><circle cx="23" cy="30" r="2" fill="#fff"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPear0)"/>`),
		g.Group(children),
	)
}