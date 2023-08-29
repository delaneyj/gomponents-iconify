package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flutter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<defs><linearGradient id="logosFlutter0" x1="3.952%" x2="75.897%" y1="26.993%" y2="52.919%"><stop offset="0%"/><stop offset="100%" stop-opacity="0"/></linearGradient></defs><path fill="#47C5FB" d="M157.666.001L.001 157.666l48.8 48.8L255.268.001zm-1.099 145.396l-84.418 84.418l48.984 49.716l48.71-48.71l85.425-85.424z"/><path fill="#00569E" d="m121.133 279.531l37.082 37.082h97.052l-85.425-85.792z"/><path fill="#00B5F8" d="m71.6 230.364l48.801-48.801l49.441 49.258l-48.709 48.71z"/><path fill="url(#logosFlutter0)" fill-opacity=".8" d="m121.133 279.531l40.56-13.459l4.029-31.131z"/>`),
		g.Group(children),
	)
}