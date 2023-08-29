package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yahoo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.818 13.803L11.005 24a8.65 8.65 0 0 0-1.505-.16h-.009a9.34 9.34 0 0 0-1.573.17l.059-.01l.187-10.197q-.577-.995-2.43-4.262t-3.12-5.402Q1.343 2.005.001 0a6.1 6.1 0 0 0 1.554.216h.004a7.6 7.6 0 0 0 1.653-.227L3.159 0q.909 1.6 1.926 3.31t2.409 3.988l2 3.274q.534-.88 1.579-2.56t1.694-2.74t1.514-2.538T15.823 0c.463.127.994.2 1.542.202h.001a6.977 6.977 0 0 0 1.692-.212L19.01 0q-.4.56-.866 1.28t-.714 1.132q-.252.418-.815 1.385t-.706 1.211q-2.106 3.574-5.091 8.795z"/>`),
		g.Group(children),
	)
}