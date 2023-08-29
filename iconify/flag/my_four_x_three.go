package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g clip-path="url(#flagMy4x30)"><path fill="#C00" d="M0 0h640v480H0V0Z"/><path fill="#C00" d="M0 0h640v34.3H0z"/><path fill="#fff" d="M0 34.3h640v34.3H0z"/><path fill="#C00" d="M0 68.6h640v34.3H0z"/><path fill="#fff" d="M0 102.9h640V137H0z"/><path fill="#C00" d="M0 137.1h640v34.3H0z"/><path fill="#fff" d="M0 171.4h640v34.3H0z"/><path fill="#C00" d="M0 205.7h640V240H0z"/><path fill="#fff" d="M0 240h640v34.3H0z"/><path fill="#C00" d="M0 274.3h640v34.3H0z"/><path fill="#fff" d="M0 308.6h640v34.3H0z"/><path fill="#C00" d="M0 342.9h640V377H0z"/><path fill="#fff" d="M0 377.1h640v34.3H0z"/><path fill="#C00" d="M0 411.4h640v34.3H0z"/><path fill="#fff" d="M0 445.7h640V480H0z"/><path fill="#006" d="M0 .5h320v274.3H0V.5Z"/><path fill="#FC0" d="m207.5 73.8l6 40.7l23-34l-12.4 39.2l35.5-20.8l-28.1 30l41-3.2l-38.3 14.8l38.3 14.8l-41-3.2l28.1 30l-35.5-20.8l12.3 39.3l-23-34.1l-6 40.7l-5.9-40.7l-23 34l12.4-39.2l-35.5 20.8l28-30l-41 3.2l38.4-14.8l-38.3-14.8l41 3.2l-28.1-30l35.5 20.8l-12.4-39.3l23 34.1l6-40.7Zm-33.3 1.7a71.1 71.1 0 0 0-100 65a71.1 71.1 0 0 0 100 65a80 80 0 0 1-83.2 6.2a80 80 0 0 1-43.4-71.2a80 80 0 0 1 126.6-65Z"/></g><defs><clipPath id="flagMy4x30"><path fill="#fff" d="M0 0h640v480H0z"/></clipPath></defs>`),
		g.Group(children),
	)
}