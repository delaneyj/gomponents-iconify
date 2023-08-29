package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Libary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.192 35.498l3.556 2.166l-1.12-16.19m-9.834 19.889h1.603v-1.069l3.206-1.068v2.137h1.602"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.406 28.202l.991 12.092c3.388-5.372 2.869-11.294 4.12-17.878"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.839 15.72c-1.925 3.241-5.234 4.187-8.954 5.853c0 0 1.431 6.628 3.521 6.63c1.793 0 2.629-3.285 2.629-3.285"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.137 20.641s.03 2.765 1.285 3.152c1.732.535 6.864-5.406 6.864-5.406l-.094 17.111l-2.59 3.728"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.315 24.267c4.444-1.16 6.476-5.1 7.846-8.548h5.678c1.84 3.814 4.171 6.818 7.846 8.548"/><ellipse cx="24" cy="13.048" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.603" ry="2.137"/><circle cx="13.315" cy="24.267" r=".379" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="34.685" cy="24.267" r=".379" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.178 43.5h25.644V17.856c0-7.486-5.28-13.35-12.822-13.356s-12.822 5.877-12.822 13.356Z"/>`),
		g.Group(children),
	)
}