package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Baloon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M412 891q76 77 95 133H261q19-56 94-133q-49-13-111.5-67t-116-125.5T37 541T0 384q0-104 51.5-192.5t140-140T384 0t192.5 51.5t140 140T768 384q0 71-37 157t-90.5 157.5t-116 125.5T412 891zM255 195q-38-11-74 16t-48 76q-14 48-5 168.5T160 576q14 0 23-32.5t18.5-73.5t22.5-54q5-5 20-19.5t23.5-23.5t20-21.5t18.5-24t9-22.5q23-87-60-110z"/>`),
		g.Group(children),
	)
}