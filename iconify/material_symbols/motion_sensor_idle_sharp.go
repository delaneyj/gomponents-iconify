package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MotionSensorIdleSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 3.825v2q-.2.05-.4.113t-.4.137l-1.5-1.5q.525-.275 1.1-.463t1.2-.287Zm7.9 7.9q-.1.625-.288 1.2t-.462 1.1l-1.5-1.5q.075-.2.138-.4t.112-.4h2ZM2 20.7v-6h2v4h4v2H2Zm18-14v-4h-4v-2h6v6h-2Zm-13.175-4l-2-2H8v2H6.825ZM22 17.875l-2-2V14.7h2v3.175Zm-16.9-6.15h2q.3 1.475 1.363 2.538T11 15.625v2q-2.3-.35-3.925-1.975T5.1 11.725Zm1.975-5.95l1.4 1.4q-.5.525-.862 1.175T7.1 9.725h-2q.175-1.15.688-2.15t1.287-1.8Zm8.475 8.475l1.4 1.4q-.8.775-1.8 1.288t-2.15.687v-2q.725-.15 1.375-.513t1.175-.862ZM13 3.825q2.3.35 3.938 1.975T18.9 9.725h-2q-.3-1.475-1.362-2.537T13 5.825v-2ZM19.15 20.7H16v-2h1.15L4 5.55V6.7H2V3.55L.675 2.225L2.1.8l19.8 19.8l-1.425 1.425L19.15 20.7Z"/>`),
		g.Group(children),
	)
}