package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Discord(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor" fill-opacity="0"><circle cx="9" cy="12" r="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.4s" values="0;1"/></circle><circle cx="15" cy="12" r="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="1.4s" dur="0.4s" values="0;1"/></circle></g><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="30" stroke-dashoffset="30" d="M15.5 17.5L16.5 19.5C16.5 19.5 20.671 18.172 22 16C22 15 22.53 7.853 19 5.5C17.5 4.5 15 4 15 4L14 6H12M8.52799 17.5L7.52799 19.5C7.52799 19.5 3.35699 18.172 2.02799 16C2.02799 15 1.49799 7.853 5.02799 5.5C6.52799 4.5 9.02799 4 9.02799 4L10.028 6H12.028"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="30;60"/></path><path stroke-dasharray="16" stroke-dashoffset="16" d="M5.5 16C10.5 18.5 13.5 18.5 18.5 16"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.4s" values="16;0"/></path></g>`),
		g.Group(children),
	)
}