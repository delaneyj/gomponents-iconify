package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonCheckmarkCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M10 8c-.792 0-1.5.679-1.5 1.6s.708 1.6 1.5 1.6s1.5-.679 1.5-1.6S10.792 8 10 8ZM6.5 9.6C6.5 7.65 8.03 6 10 6s3.5 1.65 3.5 3.6c0 1.95-1.53 3.6-3.5 3.6s-3.5-1.65-3.5-3.6Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M4 17.833c0-3.295 2.79-5.766 6.013-5.766c3.232 0 5.987 2.478 5.987 5.766V20a1 1 0 1 1-2 0v-2.167c0-2.08-1.753-3.766-3.987-3.766c-2.24 0-4.013 1.692-4.013 3.766l.002 2.166a1 1 0 0 1-2 .002L4 17.833ZM21.597 7.126a1 1 0 0 1 .388 1.36l-2.777 5a1 1 0 1 1-1.749-.972l2.778-5a1 1 0 0 1 1.36-.388Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M14.775 10.153a1 1 0 0 1 1.405-.156l2.778 2.222a1 1 0 1 1-1.25 1.562l-2.777-2.222a1 1 0 0 1-.156-1.406Z" clip-rule="evenodd"/><path d="M4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}