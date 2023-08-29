package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BitbucketOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M.5.5V0a.5.5 0 0 0-.495.57L.5.5Zm14 0l.495.07A.5.5 0 0 0 14.5 0v.5Zm-2 14v.5a.5.5 0 0 0 .495-.43l-.495-.07Zm-10 0l-.495.07A.5.5 0 0 0 2.5 15v-.5ZM5 4.5V4a.5.5 0 0 0-.498.542L5 4.5Zm4.5 6v.5a.5.5 0 0 0 .498-.459L9.5 10.5Zm-4 0l-.498.041A.5.5 0 0 0 5.5 11v-.5ZM.5 1h14V0H.5v1ZM14.005.43l-2 14l.99.14l2-14l-.99-.14ZM12.5 14h-10v1h10v-1Zm-9.505.43l-2-14l-.99.14l2 14l.99-.14ZM5 5h5V4H5v1Zm4.502-.542l-.5 6l.996.083l.5-6l-.996-.083ZM9.5 10h-4v1h4v-1Zm-3.502.459l-.5-6l-.996.083l.5 6l.996-.083ZM10 5h4V4h-4v1Z"/>`),
		g.Group(children),
	)
}