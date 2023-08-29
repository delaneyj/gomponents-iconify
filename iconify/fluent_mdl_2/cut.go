package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1416 1536q51 2 96 22t79 56t53 81t20 97q0 53-20 99t-55 81t-82 55t-99 21q-53 0-99-20t-81-55t-55-81t-21-100q0-71 36-132t100-94l-266-531l-265 529q32 17 57 41t44 54t28 63t10 70q0 53-20 99t-55 81t-82 55t-99 21q-53 0-99-20t-81-55t-55-81t-21-100q0-51 19-96t52-80t77-56t96-24l322-646l-339-678l58-175l353 708l353-708l58 175l-339 678l322 646zm-776 384q27 0 50-10t40-27t28-41t10-50q0-27-10-50t-27-40t-41-28t-50-10q-27 0-50 10t-40 27t-28 41t-10 50q0 27 10 50t27 40t41 28t50 10zm768 0q27 0 50-10t40-27t28-41t10-50q0-27-10-50t-27-40t-41-28t-50-10q-27 0-50 10t-40 27t-28 41t-10 50q0 27 10 50t27 40t41 28t50 10z"/>`),
		g.Group(children),
	)
}