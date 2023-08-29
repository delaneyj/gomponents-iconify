package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wercker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 264h35.557V132.8L191.991 40l156.46 92.8v166.4c0 25.6-9.251 46.4-25.251 68.8c-38.682 52.8-108.8 91.2-131.2 102.4c-22.044-11.2-86.4-46.4-128-96h206.209c14.94-14.4 19.391-19.2 32.191-36.8H4.943C33.848 440 182 512 182 512c-.04 0 168.4-70.4 197.2-174.4c3.014-11.2 4.8-25.6 4.8-38.4V112L191.991 0L0 112v152zm282.044-100.8v30.4H99.126v-30.4h182.918zm15.196 54.4V248H83.838v-30.4H297.24zm-3.838 52.8v30.4H80v-30.4h213.402zm-60.943-160v30.4h-83.838v-30.4h83.838z"/>`),
		g.Group(children),
	)
}