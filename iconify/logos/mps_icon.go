package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MpsIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<defs><linearGradient id="logosMpsIcon0" x1="75.7%" x2="-19.467%" y1="132.917%" y2="11.033%"><stop offset="6%" stop-color="#087CFA"/><stop offset="87%" stop-color="#21D789"/></linearGradient><linearGradient id="logosMpsIcon1" x1="76.465%" x2="28.668%" y1="92.417%" y2="-29.467%"><stop offset="5%" stop-color="#087CFA"/><stop offset="18%" stop-color="#0A84F0"/><stop offset="39%" stop-color="#1099D6"/><stop offset="67%" stop-color="#19BAAD"/><stop offset="87%" stop-color="#21D789"/></linearGradient><linearGradient id="logosMpsIcon2" x1="33.801%" x2="64.854%" y1="103.028%" y2="-21.398%"><stop offset="12%" stop-color="#21D789"/><stop offset="36%" stop-color="#6AE274"/><stop offset="58%" stop-color="#A9EB62"/><stop offset="77%" stop-color="#D6F255"/><stop offset="92%" stop-color="#F2F64D"/><stop offset="100%" stop-color="#FCF84A"/></linearGradient></defs><path fill="url(#logosMpsIcon0)" d="M0 256h256L127.403 127.659L0 0z"/><path fill="url(#logosMpsIcon1)" d="M256 256L127.403 127.659L256 0z"/><path fill="url(#logosMpsIcon2)" d="m191.573 191.616l-64.17-63.957L256 0z"/>`),
		g.Group(children),
	)
}