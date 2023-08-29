package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ptcltouch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.063 28.023c-1.322 7.879-10.418 10.965-16.847 7.63a15.548 15.548 0 0 1-7.32-8.37a18.917 18.917 0 0 1-1.138-7.336q.017-.378.048-.756a18.833 18.833 0 0 1 .416-2.69a17.396 17.396 0 0 1 8.958-11.61a19.92 19.92 0 0 1 14.368-1.357a20.76 20.76 0 0 1 12.378 9.851q.248.439.476.887c4.088 8.36 2.247 20.657-5.926 26.645c-6.222 4.428-15.368 5.75-22.386 2.48c-4.564-1.873-9.341-7.529-11.299-13.083a29.447 29.447 0 0 1-.968-13.07"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.005 19.02a4.305 4.305 0 1 1-4.305 4.304a4.305 4.305 0 0 1 4.305-4.304Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.063 28.023c-1.39 4.156-5.683 6.91-10.307 6.91a10.477 10.477 0 1 1 3.223-20.449"/>`),
		g.Group(children),
	)
}