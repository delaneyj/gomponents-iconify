package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cinema(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#1b75bb" d="M64 57.1a6.9 6.9 0 0 1-6.898 6.904H6.892c-3.808 0-6.894-3.091-6.894-6.904V6.9c0-3.81 3.086-6.9 6.894-6.9h50.21A6.899 6.899 0 0 1 64 6.9v50.2"/><g fill="#fff"><path d="M9.06 27.369c-1.686.131-3.091 1.01-3.706 2.201v23.855c.617 1.198 2.028 2.073 3.717 2.207l16.3-14.08L9.06 27.373"/><path d="M53.09 46.806c0 5.424-3.835 9.819-8.562 9.819h-14.09c-4.726 0-8.561-4.396-8.561-9.819V36.19c0-5.424 3.835-9.819 8.561-9.819h14.09c4.727 0 8.562 4.396 8.562 9.819v10.616"/><circle cx="25.374" cy="16.186" r="9.215"/><circle cx="49.592" cy="16.186" r="9.213"/></g>`),
		g.Group(children),
	)
}