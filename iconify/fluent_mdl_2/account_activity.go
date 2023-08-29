package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccountActivity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 256v1792H256V256h512q0-53 20-99t55-82t81-55t100-20q53 0 99 20t82 55t55 81t20 100h512zM640 384v128h768V384h-256V256q0-27-10-50t-27-40t-41-28t-50-10q-27 0-50 10t-40 27t-28 41t-10 50v128H640zm1024 0h-128v256H512V384H384v1536h1280V384zm-405 335q42 0 78 15t64 42t42 63t16 78q0 39-15 76t-43 65l-526 531l-358 68l75-351l526-530q28-28 65-42t76-15zm51 249q21-23 21-51q0-31-20-50t-52-20q-14 0-27 4t-23 15l-499 503l-27 126l129-25l498-502z"/>`),
		g.Group(children),
	)
}