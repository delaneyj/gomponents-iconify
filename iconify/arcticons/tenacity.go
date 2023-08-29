package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tenacity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" d="M13.623 43.5a18.004 18.004 0 0 1-6.395-21.71a16.338 16.338 0 0 1 9.485-9.2a13.988 13.988 0 0 1 6.712-.617a12.836 12.836 0 0 1 6.144 2.74a11.787 11.787 0 0 1 3.565 4.994a10.855 10.855 0 0 1 .412 6.109a10.063 10.063 0 0 1-3.096 5.266a9.38 9.38 0 0 1-5.626 2.342a8.368 8.368 0 0 1-6.416-2.413a7.382 7.382 0 0 1-1.779-3.015a6.885 6.885 0 0 1-.12-3.493a6.31 6.31 0 0 1 1.171-2.474a5.733 5.733 0 0 1 2.146-1.691a5.257 5.257 0 0 1 2.692-.44a4.852 4.852 0 0 1 2.52 1.032a4.64 4.64 0 0 1 1.116 5.911"/><path fill="none" stroke="currentColor" stroke-linecap="round" d="M10.597 6.498a19.203 19.203 0 0 1 17.421.186a19.2 19.2 0 0 1 9.846 14.373a16.925 16.925 0 0 1-1.25 8.896a13.944 13.944 0 0 1-6.012 6.603a12.496 12.496 0 0 1-9.803.855a11.839 11.839 0 0 1-7.129-6.736a10.799 10.799 0 0 1-.592-5.949a9.725 9.725 0 0 1 2.842-5.238a8.584 8.584 0 0 1 6.05-2.359a7.467 7.467 0 0 1 5.728 2.942a6.694 6.694 0 0 1 1.217 2.9a6.282 6.282 0 0 1-.298 3.127a5.92 5.92 0 0 1-1.82 2.555a5.595 5.595 0 0 1-2.889 1.21a4.62 4.62 0 0 1-3.472-.988a3.942 3.942 0 0 1-.16-5.858"/><path fill="none" stroke="currentColor" stroke-linecap="round" d="M41.356 10.356a32.438 32.438 0 0 1-.09 14.535a21.201 21.201 0 0 1-1.408 4.175a15.757 15.757 0 0 1-15.62 8.964"/><path fill="none" stroke="currentColor" stroke-linecap="round" d="M34.518 8.3a20.771 20.771 0 0 1 5.071 21.307"/><ellipse cx="23.154" cy="25.02" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.531" ry="3.396"/>`),
		g.Group(children),
	)
}