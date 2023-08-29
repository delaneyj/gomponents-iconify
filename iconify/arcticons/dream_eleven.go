package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DreamEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.71 33.14c4.823-3.645 9.168-8.44 9.168-13.037v-9.224C32.889 10.094 29.72 9.028 24 9.028s-8.888 1.065-10.879 1.85v9.225c0 4.598 4.346 9.392 9.169 13.037v1.01c-2.524 1.261-3.925 3.28-3.925 3.28v1.542h11.271V37.43s-1.13-1.626-3.185-2.871c-2.03 1.02-5.283 2.619-5.283 2.619"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.29 33.14C8.663 24.981 9.252 23.243 9.252 15.673c.73-.701 2.3-1.29 2.3-1.29M25.71 33.14c13.627-8.159 13.038-9.897 13.038-17.467c-.73-.701-2.3-1.29-2.3-1.29M21.557 25.13V12.87h3.364a3.59 3.59 0 0 1 3.59 3.59v5.08a3.59 3.59 0 0 1-3.59 3.59h-3.364Zm0-12.26h-2.068m2.068 12.26h-2.068"/>`),
		g.Group(children),
	)
}