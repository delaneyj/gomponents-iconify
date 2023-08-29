package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Evmap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.597 23.17a12.446 12.446 0 1 0-21.558 0s9.045 12.927 10.779 20.33c1.358-7.37 10.779-20.33 10.779-20.33Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.706 11.347h7.053l-2.593 5.186h2.593l-4.978 7.779v-5.704h-2.075Zm-1.034 22.515a10.331 10.331 0 0 0-2.106 2.95c-.287 1.031.737 2.254.239 3.2c-.99 1.88-3.478 3.825-5.509 3.203c-.896-.274-1.313-2.637-1.313-2.637"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.574 40.688l2.885-.227l-.225-2.627l1.943-1.63l-.28-2.8l1.035-.12l-.174-2.12l-9.023.628l.148 2.326l.97-.035l.3 2.758l2.174 1.317Zm-1.729-12.585l.236 3.524m3.643-3.825l.27 3.534"/>`),
		g.Group(children),
	)
}