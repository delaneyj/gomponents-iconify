package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditStyle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1055 896H609l-85 256H384L768 0h128l330 988l-106 105l-65-197zm-43-128L832 228L652 768h360zm581 131q42 0 78 14t63 41t42 61t16 79q0 39-15 76t-43 65l-717 717l-377 94l94-377l717-716q28-28 65-41t77-13zm50 246q21-21 21-51q0-32-20-50t-52-19q-14 0-27 4t-23 14l-692 692l-34 135l135-34l692-691z"/>`),
		g.Group(children),
	)
}