package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DjMan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<rect width="40" height="6.677" x="16" y="49.322" fill="#9b9b9a" rx="3.339" ry="3.339"/><path fill="#d0cfce" d="M56 52.661V56H39.924v-6.677H56v3.338"/><g fill="#fcea2b"><circle cx="35.998" cy="19.886" r="3.5"/><path d="M53.65 15.03a1.304 1.304 0 0 0-1.683.89l-2.318 7.638a2.714 2.714 0 0 1-1.607 1.322l-3.747.919c-1.29.325-2.597.57-3.917.732l-8.694.644a5.882 5.882 0 0 0-3.721 1.154l-.62.461l-.15.127a10.186 10.186 0 0 0-2.41 3.477l-.686 2.05c-.218.617-.551 1.55-.872 2.334a4.67 4.67 0 0 0 .05 3.823l3.803 5.048c.244.537.779.882 1.368.882c.206 0 .043.042.244-.048a1.344 1.344 0 0 0 .749-1.84l-3.11-5.557a1.71 1.71 0 0 1 .018-1.399c.481-.964 1.018-1.9 1.607-2.802l1.051 8.052a.974.974 0 0 0 .965.848l10.979.013c.565 0 1.024-.458 1.025-1.023v-6.947a10.61 10.61 0 0 1 .936-4.681l6.314-3.071a6.164 6.164 0 0 0 3.274-3.577l2.02-7.644a1.501 1.501 0 0 0-.868-1.824Z"/></g><circle cx="36.002" cy="19.886" r="3.5" fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2.333"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m41.971 30.792l-.5 1.5a13.403 13.403 0 0 0-.5 3.5v7m-11 0l-1-8"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m27.657 45.649l-4.509-5.66a3.353 3.353 0 0 1-.109-2.708a60.1 60.1 0 0 0 .892-2.377l.765-2.22a10.338 10.338 0 0 1 2.067-3.306l.172-.171a6.29 6.29 0 0 1 3.408-1.568l9.018-.694a28.639 28.639 0 0 0 3.928-.638l4.12-1.03a4.138 4.138 0 0 0 2.596-2.373l2.469-7.111"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.216" d="M30.745 20.455v-3.236a5.262 5.262 0 0 1 5.262-5.262h0a5.262 5.262 0 0 1 5.263 5.262v3.236"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21.534 46.483H34.3m3.393 0h12.765m5.54 6.178v-3.338h-40V56h40v-3.339z"/><circle cx="46.657" cy="52.661" r="1"/><circle cx="49.645" cy="52.661" r="1"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.216" d="M30.526 20.61a1.634 1.634 0 1 1 0-3.269m10.731 0a1.634 1.634 0 1 1 0 3.269"/>`),
		g.Group(children),
	)
}