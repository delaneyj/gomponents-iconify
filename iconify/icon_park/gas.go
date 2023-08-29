package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gas(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path fill="#2F88FF" stroke-linejoin="round" d="M23.0488 9.78854C24.6746 9.35291 26.3402 10.2971 26.769 11.8975C27.1978 13.4979 26.2275 15.1485 24.6017 15.5841C22.9759 16.0197 6.70807 17 6.70807 17C6.70807 17 21.423 10.2242 23.0488 9.78854Z"/><path fill="#2F88FF" stroke-linejoin="round" d="M23.1066 38.4467C24.7324 38.8823 26.398 37.9381 26.8268 36.3377C27.2556 34.7373 26.2853 33.0868 24.6595 32.6511C23.0337 32.2155 6.7659 31.0039 6.7659 31.0039C6.7659 31.0039 21.4808 38.011 23.1066 38.4467Z"/><path stroke-linecap="round" d="M33.9996 16.0039C34.9118 14.7895 36.3642 14.0039 38 14.0039C40.7614 14.0039 43 16.2425 43 19.0039C43 21.7653 40.7614 24.0039 38 24.0039H16"/></g>`),
		g.Group(children),
	)
}