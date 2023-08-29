package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MjPdfReader(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.295 10.402h5.59v8.384a2.803 2.803 0 0 1-2.795 2.795h0a2.803 2.803 0 0 1-2.795-2.795v-.978m-14.179 3.773V10.402l5.589 11.179l5.59-11.179v11.179m-4.171 16.018v-11.18h1.817c3.074 0 5.59 2.515 5.59 5.59h0c0 3.074-2.516 5.59-5.59 5.59h-1.817Zm10.786-11.18h5.59m-5.59 5.59h3.634m-3.634-5.59v11.18m-21.41 0v-11.18h3.633c2.097 0 3.774 1.677 3.774 3.773s-1.677 3.773-3.774 3.773H10.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}