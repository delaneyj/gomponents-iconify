package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunriseOverMountains(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#e27022" d="M4 4h64v64H4z"/><path fill="#f4aa41" d="M68 4h-8L44 35l4 3l17-20l3-3.231M4 17.562l2.019 1.913L24.55 38.066l3.75-3.308L9 4H4v13.562zM30.879 4l2.469 28.375l4.983.412l3.749-25.98L42.566 4"/><path fill="#fcea2b" d="M50.961 49.078c-.267 3.131-5.203 14.664-16.04 13.883a15 15 0 0 1-13.882-16.04c.627-8.7 8.662-14.341 16.04-13.882c6.991.435 11.089 6.206 11.654 7.03a14.632 14.632 0 0 1 2.107 5.547a16.402 16.402 0 0 1 .121 3.462Z"/><path fill="#d0cfce" d="m4 42.636l11.889-9.226L54.869 68H4V42.636z"/><path fill="#9b9b9a" d="M41.492 54.869L68 27.944V68H54.869"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-width="2"><path stroke-miterlimit="10" d="M25.75 37.049a15.025 15.025 0 0 1 14.712-3.374a15.362 15.362 0 0 1 7.586 5.388a11.518 11.518 0 0 1 1.694 2.88"/><path stroke-linejoin="round" d="m5 42.416l11.34-9.561L54.746 67M42.983 53.445L66.755 28.22"/></g>`),
		g.Group(children),
	)
}