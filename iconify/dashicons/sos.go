package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 10c0-4.42-3.58-8-8-8s-8 3.58-8 8s3.58 8 8 8s8-3.58 8-8zM7.23 3.57L8.72 7.3c-.62.29-1.13.8-1.42 1.42L3.57 7.23a7.037 7.037 0 0 1 3.66-3.66zm9.2 3.66L12.7 8.72c-.29-.62-.8-1.13-1.42-1.42l1.49-3.73c1.64.71 2.95 2.02 3.66 3.66zM10 12c-1.1 0-2-.9-2-2s.9-2 2-2s2 .9 2 2s-.9 2-2 2zm-6.43.77l3.73-1.49c.29.62.8 1.13 1.42 1.42l-1.49 3.73a7.037 7.037 0 0 1-3.66-3.66zm9.2 3.66l-1.49-3.73c.62-.29 1.13-.8 1.42-1.42l3.73 1.49a7.037 7.037 0 0 1-3.66 3.66z"/>`),
		g.Group(children),
	)
}