package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Station(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#009E83" d="M240.019 64H16v368h320V64h-95.981zM176 343.019l-152-152V72h95.981L176 128.019L232.019 72H328v119.019l-152 152z"/><path fill="#FFD469" d="M128 340c0 11.046-8.954 20-20 20s-20-8.954-20-20s8.954-20 20-20s20 8.954 20 20zm116-20c-11.046 0-20 8.954-20 20s8.954 20 20 20s20-8.954 20-20s-8.954-20-20-20zm-180 8c-6.627 0-12 5.373-12 12s5.373 12 12 12s12-5.373 12-12s-5.373-12-12-12zm222 0c-6.627 0-12 5.373-12 12s5.373 12 12 12s12-5.373 12-12s-5.373-12-12-12z"/><path fill="#575A5B" d="M16 384h320v48H16z"/><path fill="#132028" d="m49 496l64-64h32l-16 64H49zm255 0l-64-64h-32l16 64h80zM88 128H36v96h52v-96zm8 96h160v-96H96v96zm168-96v96h52v-96h-52zM176 16v32h288v384H352v64h144V16H176z"/>`),
		g.Group(children),
	)
}