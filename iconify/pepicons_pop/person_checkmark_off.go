package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonCheckmarkOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M7 5c-.792 0-1.5.679-1.5 1.6S6.208 8.2 7 8.2s1.5-.679 1.5-1.6S7.792 5 7 5ZM3.5 6.6C3.5 4.65 5.03 3 7 3s3.5 1.65 3.5 3.6c0 1.95-1.53 3.6-3.5 3.6S3.5 8.55 3.5 6.6Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M1 14.833c0-3.295 2.79-5.766 6.013-5.766c3.232 0 5.987 2.478 5.987 5.766V17a1 1 0 1 1-2 0v-2.167c0-2.08-1.753-3.766-3.987-3.766c-2.24 0-4.013 1.692-4.013 3.766l.002 2.166a1 1 0 0 1-2 .002L1 14.833ZM18.597 4.126a1 1 0 0 1 .388 1.36l-2.777 5a1 1 0 1 1-1.749-.972l2.778-5a1 1 0 0 1 1.36-.388Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M11.775 7.153a1 1 0 0 1 1.405-.156l2.778 2.222a1 1 0 1 1-1.25 1.562l-2.777-2.222a1 1 0 0 1-.156-1.406Z" clip-rule="evenodd"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}