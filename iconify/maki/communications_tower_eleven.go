package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommunicationsTowerEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M7.655 3.564h-.4V1.448h.4zM3.702 1.448h-.4v2.116h.4zm6-.285h-.4v3.278h.4zm-8.023 0h-.4v3.278h.4zm4.078 3.005h.437a.354.354 0 0 1 .346.282l1.04 4.987H9V10H2v-.563h1.434l1.04-4.987a.355.355 0 0 1 .344-.281h.44V2.912a.915.915 0 0 1-.668-.875a.927.927 0 1 1 1.167.888zm-.96 2.188h1.419l-.309-1.481h-.801zm-.414 1.989H6.63l-.351-1.688H4.735zm2.475 1.092l-.165-.791H4.32l-.165.79z" fill="currentColor"/>`),
		g.Group(children),
	)
}