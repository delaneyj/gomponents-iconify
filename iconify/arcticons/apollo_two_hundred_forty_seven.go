package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApolloTwoHundredFortySeven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.71 16.716h0a2.046 2.046 0 0 1-2.04-2.04V13.35c0-1.122.919-2.04 2.04-2.04h0c1.123 0 2.04.918 2.04 2.04v1.326c0 1.122-.917 2.04-2.04 2.04Zm-4.952-8.16v8.16m1.481-8.16v8.16m-4.96 0h0a2.046 2.046 0 0 1-2.04-2.04V13.35c0-1.122.918-2.04 2.04-2.04h0c1.122 0 2.04.918 2.04 2.04v1.326c0 1.122-.918 2.04-2.04 2.04Zm-7.766-1.989c0 1.122.918 2.04 2.04 2.04h0c1.122 0 2.04-.918 2.04-2.04V13.4c0-1.122-.918-2.04-2.04-2.04h0c-1.122 0-2.04.918-2.04 2.04m0-2.039v8.16m-1.607-2.805l-2.652-8.16l-2.754 8.16m.918-2.754h3.57M35.689 32.24l4.691-8.727h-5.782m-6.152 8.727v-8.727l-4.691 5.891h5.782m-12.977-3c0-1.636 1.31-2.945 2.837-2.945s2.945 1.309 2.945 2.945a2.98 2.98 0 0 1-.872 2.073c-1.2.982-4.91 3.818-4.91 3.818h5.782M32.33 21.41l.074 12.933"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.48 18.91h27.956c.59 0 1.064.474 1.064 1.064v15.683c0 .59-.474 1.064-1.064 1.064H18.171c-2.962 0-4.755 3.24-4.755 2.652v-19.4c0-.588.474-1.063 1.064-1.063Z"/>`),
		g.Group(children),
	)
}