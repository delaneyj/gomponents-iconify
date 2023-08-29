package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonMountainBiking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d0cfce" d="M40 26.095L56 16l12 13.46V69H4V29.46l14-8.412l10 6.73l12-1.683z"/><circle cx="19.1" cy="57" r="10.2" fill="#fff"/><circle cx="51.1" cy="57" r="10.2" fill="#fff"/><g fill="#fcea2b"><circle cx="33.4" cy="16.1" r="3"/><path d="m33.4 26.9l2-2.9l3.8.1l6.3 14.1l-2.1 3.8l-2.5 2.5l-3.1 11.2l-3.4 6.7l-2.4-2l.9-7.1l-3.8-7l.1-3.3l8-5l-2.8-7.5"/></g><g fill="none" stroke="#000" stroke-width="2"><circle cx="33.4" cy="17.1" r="3" stroke-miterlimit="10"/><path stroke-linecap="round" stroke-linejoin="round" d="m25.3 34.1l4.3-2.9a12.399 12.399 0 0 0 2.8-2.8l1.9-2.7a4.135 4.135 0 0 1 3.1-1.6a3.205 3.205 0 0 1 2.7 1.9l4.6 10.5a4.974 4.974 0 0 1-.1 3.6l-.1.2a6.163 6.163 0 0 1-2.8 2.5l-5.5 2.4a2.03 2.03 0 0 0-1.1 2.6l2.6 6.4A2.397 2.397 0 0 1 37 57a1.815 1.815 0 0 1-2.5-.8L30.1 48a9.073 9.073 0 0 1-.9-3.4a3.846 3.846 0 0 1 1.7-2.7l4.6-2.9a2.92 2.92 0 0 0 1.2-3l-1.5-6"/><circle cx="19.1" cy="58" r="10.2" stroke-miterlimit="10"/><circle cx="51.1" cy="58" r="10.2" stroke-miterlimit="10"/><path stroke-linecap="round" stroke-linejoin="round" d="m19.2 58l2-24H25"/><path stroke-miterlimit="10" d="m37 57l-1 3.1a2.791 2.791 0 0 1-2.4 1.9a1.702 1.702 0 0 1-1.6-2l.9-6.7m8.5-10.4l-3.6 11.5"/><path stroke-linecap="round" stroke-miterlimit="10" d="m5 29.46l10.362-7.265a3.262 3.262 0 0 1 3.33-.08L25 26.095m19-2.524l10.309-6.504a2.326 2.326 0 0 1 2.957.482L67 29.46"/></g>`),
		g.Group(children),
	)
}