package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IlFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagIl4x30"><path fill-opacity=".7" d="M-87.6 0H595v512H-87.6z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagIl4x30)" transform="translate(82.1) scale(.94)"><path fill="#fff" d="M619.4 512H-112V0h731.4z"/><path fill="#0038b8" d="M619.4 115.2H-112V48h731.4zm0 350.5H-112v-67.2h731.4zm-483-275l110.1 191.6L359 191.6l-222.6-.8z"/><path fill="#fff" d="m225.8 317.8l20.9 35.5l21.4-35.3l-42.4-.2z"/><path fill="#0038b8" d="M136 320.6L246.2 129l112.4 190.8l-222.6.8z"/><path fill="#fff" d="m225.8 191.6l20.9-35.5l21.4 35.4l-42.4.1zM182 271.1l-21.7 36l41-.1l-19.3-36zm-21.3-66.5l41.2.3l-19.8 36.3l-21.4-36.6zm151.2 67l20.9 35.5l-41.7-.5l20.8-35zm20.5-67l-41.2.3l19.8 36.3l21.4-36.6zm-114.3 0L189.7 256l28.8 50.3l52.8 1.2l32-51.5l-29.6-52l-55.6.5z"/></g>`),
		g.Group(children),
	)
}