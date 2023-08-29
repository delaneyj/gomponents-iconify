package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeShowTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M1.974 6.659a.5.5 0 0 1-.948-.317c-.01.03 0-.001 0-.001a1.661 1.661 0 0 1 .062-.162c.04-.095.099-.226.18-.381c.164-.31.422-.723.8-1.136C2.835 3.827 4.088 3 6 3c1.913 0 3.166.827 3.931 1.662a5.473 5.473 0 0 1 .98 1.517c.02.047.035.085.045.113c.004.008.014.06.024.11c.009.046.017.09.02.098c0 0 .084.333-.342.474a.5.5 0 0 1-.632-.314v-.003l-.007-.016a3.676 3.676 0 0 0-.172-.376a4.477 4.477 0 0 0-.653-.927C8.584 4.673 7.587 4 6 4s-2.584.673-3.194 1.338a4.477 4.477 0 0 0-.795 1.225a2.184 2.184 0 0 0-.03.078l-.007.018z" fill="currentColor"/><path d="M6 5a2 2 0 1 0 0 4a2 2 0 0 0 0-4zM5 7a1 1 0 1 1 2 0a1 1 0 0 1-2 0z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}