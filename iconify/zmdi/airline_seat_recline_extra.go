package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatReclineExtra(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M71.5 80Q57 70 54 52.5t7-32T88.5 3t32 7.5T138 38t-7 31.5T103.5 87t-32-7zM299 365v43H148q-39 0-69-25.5T42 318L0 109h43l42 202q3 24 21 39t42 15h151zm5-85l123 96l-32 32l-82-64H167q-23 0-40.5-14.5T104 292L75 166q-3-20 8.5-36.5T115 110q10-2 21 1q10 3 16 8l35 27q47 37 100 27v46q-48 8-110-26l22 87h105z"/>`),
		g.Group(children),
	)
}