package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForLiberia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#e6e7e8" d="M0 30h64v4H0z"/><path fill="#ec1c24" d="M0 34h64v4H0z"/><path fill="#e6e7e8" d="M0 38h64v4H0z"/><path fill="#ec1c24" d="M0 44c0 .684.049 1.351.135 2h63.73c.086-.649.135-1.316.135-2v-2H0v2"/><path fill="#e6e7e8" d="M1.346 50h61.31c.616-1.193 1.021-2.546 1.211-4H.136c.19 1.454.594 2.807 1.211 4"/><path fill="#ec1c24" d="M62.65 50H1.34c.873 2.063 4.443 4 8.654 4h44c4.211 0 7.725-2.023 8.654-4m-.578-36c-1.647-2.443-4.345-4-8.07-4H27v4h35.07"/><path fill="#e6e7e8" d="M27 18h36.695c-.307-1.483-.844-2.843-1.624-4h-35.07v4"/><path fill="#fff" d="M63.982 21c.015.365.021.707.018 1v-1h-.018"/><path fill="#e6e7e8" d="M27 22h37v4H27z"/><path fill="#ec1c24" d="M27 26h37v4H27zm36.982-5a28.73 28.73 0 0 0-.287-3H27v4h37c.004-.293-.003-.635-.018-1"/><path fill="#003893" d="M10 10c-3.727 0-6.424 1.557-8.07 4c-.78 1.157-1.317 2.517-1.624 4c-.143.91-.261 2.051-.299 3H.001v.123c-.011.324-.012.623 0 .877v12h27V10h-17"/><path fill="#e6e7e8" d="m23.24 19.547l-7.16.006l-2.218-7.235l-2.207 7.235l-7.155-.006l5.799 4.402L8.05 31.14l5.824-4.471l5.823 4.471l-2.257-7.191z"/>`),
		g.Group(children),
	)
}