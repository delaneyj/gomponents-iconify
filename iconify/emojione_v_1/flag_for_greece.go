package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForGreece(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#194992" d="M10 54h44c4.211 0 7.102-1.992 8.654-5H1.344c1.552 3.01 4.443 5 8.654 5M62.65 15c-1.553-3.01-4.443-5-8.654-5h-29v5H62.65"/><path fill="#e6e7e8" d="M1.346 49h61.31c.755-1.462 1.176-3.169 1.294-5H.053c.118 1.831.539 3.538 1.294 5M25 20h38.948c-.118-1.831-.539-3.538-1.294-5H25v5zm0 9h39v-4.154l-39 .039zM0 34v5h64v-5H25z"/><path fill="#194992" d="M63.948 20H25v5h39v-4c0-.34-.031-.668-.052-1M0 43c0 .34.031.668.052 1h63.896c.021-.332.052-.66.052-1v-4H0v4m25-14h39v5H25zM10 10c-4.211 0-7.102 1.992-8.654 5C.591 16.462.17 18.169.052 20c-.021.332-.052.66-.052 1v13h25V10H10"/><path fill="#e6e7e8" d="M15 10h-5v10H.052c-.021.332-.052.66-.052 1v4h10v9h5v-9h10v-5H15V10z"/>`),
		g.Group(children),
	)
}