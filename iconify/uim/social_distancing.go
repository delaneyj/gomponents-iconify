package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialDistancing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 11a3.5 3.5 0 1 1 3.5-3.5A3.504 3.504 0 0 1 6 11Z" opacity=".25"/><path fill="currentColor" d="M8.64 9.772a3.452 3.452 0 0 1-5.28 0A4.988 4.988 0 0 0 1 14a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1a4.988 4.988 0 0 0-2.36-4.228zM18 11a3.5 3.5 0 1 1 3.5-3.5A3.504 3.504 0 0 1 18 11z" opacity=".5"/><path fill="currentColor" d="M20.64 9.772a3.452 3.452 0 0 1-5.28 0A4.988 4.988 0 0 0 13 14a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1a4.988 4.988 0 0 0-2.36-4.228Z" opacity=".7"/><path fill="currentColor" d="m21.207 18.293l-2-2a1 1 0 0 0-1.414 1.414l.293.293H15.5a1 1 0 0 0 0 2h2.586l-.293.293a1 1 0 1 0 1.414 1.414l2-2a1 1 0 0 0 0-1.414zM8.5 18H5.914l.293-.293a1 1 0 0 0-1.414-1.414l-2 2a1 1 0 0 0 0 1.414l2 2a1 1 0 0 0 1.414-1.414L5.914 20H8.5a1 1 0 0 0 0-2z"/>`),
		g.Group(children),
	)
}