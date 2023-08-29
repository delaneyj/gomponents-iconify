package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudWithLightning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#fbb11c" d="M38.751 10.941L14.79 38.031h11.982L14.79 60.948l34.38-25l-17.189-6.25z"/><path fill="#bbbbbc" fill-rule="evenodd" d="M8.462 17.277c-1.131-5.959 1.139-10.48 6.309-13.253c6.911-3.708 13.59-2.881 20.04 2c3.463-3.918 7.676-5.639 12.624-3.062c4.284 2.229 5.607 6.157 5.04 10.974c1.088.097 2.072.146 3.046.296c6.967 1.052 10.781 9.569 6.88 15.564c-.488.746-1.738 1.387-2.628 1.391c-18.783.079-37.567.044-56.35.099c-1.714.003-2.462-.765-2.979-2.232c-1.028-2.914-.28-6.161 2.364-8.324c1.603-1.317 3.562-2.2 5.654-3.453"/>`),
		g.Group(children),
	)
}