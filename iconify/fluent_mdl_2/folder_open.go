package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1536 0q26 0 49 10t41 27t28 41t10 50v928q0 31 9 54t24 43t31 41t31 46t23 58t10 78v288q0 26-10 49t-27 41t-41 28t-50 10H768v128q0 39-21 71t-58 47q-23 10-49 10q-53 0-91-38l-293-293V0h1280zM640 1568q0-45 9-77t24-58t31-46t31-40t23-44t10-55V603L384 219v1445l256 256v-352zm1024-192q0-31-9-54t-24-43t-31-41t-31-46t-23-58t-10-78V128H475l384 384q18 18 27 41t10 50v645q0 45-9 77t-24 58t-31 46t-31 40t-23 44t-10 55v96h896v-288z"/>`),
		g.Group(children),
	)
}