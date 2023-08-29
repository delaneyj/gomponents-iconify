package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.127 15.608a5.217 5.217 0 0 0 5.211-5.211V5.911H1.916v4.486a5.218 5.218 0 0 0 5.21 5.211zM2.574 6.507h1.84v.406h-1.84zm0 1.475h1.84v.406h-1.84zm0 1.468h1.84v.406h-1.84zm0 1.475h1.84v.406h-1.84z"/><path fill="currentColor" d="M12.978 0H0v10.4a7.142 7.142 0 0 0 5.73 6.992l.046.008v.27c0 .21.227.375.51.375h.43v1.456h-.161a.277.277 0 0 0-.276.276v2.152c0 .152.123.274.274.274h.162V24h.83v-1.794h.162a.277.277 0 0 0 .276-.276v-2.155a.277.277 0 0 0-.276-.276h-.162v-1.456h.43c.286 0 .511-.166.511-.375v-.27c3.311-.661 5.772-3.542 5.776-6.999v-10.4zm0 10.4a5.851 5.851 0 0 1-11.7.004V1.28h11.7z"/><path fill="currentColor" d="M2.575 2.09h1.837v.406H2.575zm0 1.474h1.837v.406H2.575zm0 1.475h1.837v.406H2.575z"/>`),
		g.Group(children),
	)
}