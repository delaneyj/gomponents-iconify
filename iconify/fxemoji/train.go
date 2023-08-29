package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Train(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#009E83" d="M96 64h320v368H96z"/><path fill="#F9F9F7" d="M416 64H96v119h.019l-.019.019l160 160l160-160l-.019-.019H416z"/><path fill="#FFD469" d="M208 340c0 11.046-8.954 20-20 20s-20-8.954-20-20s8.954-20 20-20s20 8.954 20 20zm116-20c-11.046 0-20 8.954-20 20s8.954 20 20 20s20-8.954 20-20s-8.954-20-20-20zm-180 8c-6.627 0-12 5.373-12 12s5.373 12 12 12s12-5.373 12-12s-5.373-12-12-12zm222 0c-6.627 0-12 5.373-12 12s5.373 12 12 12s12-5.373 12-12s-5.373-12-12-12z"/><path fill="#132028" d="m129 496l64-64h32l-16 64h-80zm255 0l-64-64h-32l16 64h80zM168 128h-52v96h52v-96zm8 96h160v-96H176v96zm168-96v96h52v-96h-52z"/><path fill="#009E83" d="M320.019 64H96v368h320V64h-95.981zM408 424H104V72h95.981L256 128.019L312.019 72H408v352z"/><path fill="#575A5B" d="M96 384h320v48H96z"/>`),
		g.Group(children),
	)
}