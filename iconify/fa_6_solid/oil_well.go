package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OilWell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M528.3 61.3c-11.4-42.7-55.3-68-98-56.6l-15.4 4.1C397.8 13.4 387.7 31 392.3 48l24.5 91.4l-108.3 28.1l-6.3-18.1c-4.5-12.8-16.6-21.4-30.2-21.4s-25.7 8.6-30.2 21.4l-13.6 39L96 222.6V184c0-13.3-10.7-24-24-24s-24 10.7-24 24v264H32c-17.7 0-32 14.3-32 32s14.3 32 32 32h512c17.7 0 32-14.3 32-32s-14.3-32-32-32H406.7L340 257.5l-62.2 16.1l27.5 78.4h-66.6l26.3-75l-74.6 19.3L137.3 448H96V288.8l337.4-87.5l25.2 94c4.6 17.1 22.1 27.2 39.2 22.6l15.5-4.1c42.7-11.4 68-55.3 56.6-98L528.3 61.3zM205.1 448l11.2-32h111.4l11.2 32H205.1z"/>`),
		g.Group(children),
	)
}