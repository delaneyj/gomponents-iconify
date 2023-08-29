package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleFamilyLink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34.857 23.21l2.962-4.157a1.495 1.495 0 0 0-.188-1.952L25.023 5.117a2.243 2.243 0 0 0-3.09 0L9.325 17.101a1.495 1.495 0 0 0-.188 1.951l12.514 17.566a2.243 2.243 0 0 0 3.654 0l1.721-2.416"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.887 18.441l14.591-1.331l14.591 1.331m-22.986 8.958l8.395-.668l3.414.272m14.376 2.3l-7.481-7.11a1.33 1.33 0 0 0-1.834 0l-7.48 7.11a.887.887 0 0 0-.112 1.158l7.425 10.422a1.33 1.33 0 0 0 2.168 0l7.425-10.422a.887.887 0 0 0-.111-1.158Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24.213 30.098l8.657-.79l8.658.79m-13.639 5.315l4.981-.396l4.981.396m-3.897 5.47a9.323 9.323 0 0 1-4.923 2.617l.21-1.551l-1.352-1.279a11.097 11.097 0 0 0 3.033-1m-5.617-3.052c-.42.589-1.66 2.438-2.562 2.438h-6.079c-1.278 0-6.141-7.797-7.805-7.797a2.575 2.575 0 0 0-2.403 1.32l2.243.608l1.266 1.663l1.347-1.598"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.956 39.056c-1.279 0-6.142-7.797-7.805-7.797H8.86"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.652 35.215a6.09 6.09 0 0 1-1.094 1.61"/>`),
		g.Group(children),
	)
}