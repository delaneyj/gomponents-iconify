package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SchwbischGmndFnfknopfturm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#ea5a47" d="M37.116 11.241l4.417 5.809l4.417 5.81H28.282l4.417-5.81l4.417-5.809z"/><path fill="#d22f27" d="M46.949 18.484l-2.871-.049l-2.871-.049l1.478-2.513l1.478-2.514l1.393 2.562l1.393 2.563z"/><path fill="#d0cfce" d="M40.602 18.764l3.268 2.38l.08 1.688l2.098-.101l-.202-4.221"/><path fill="#d0cfce" stroke="#d0cfce" stroke-miterlimit="10" stroke-width="2" d="M39.167 13.353h-3.334L35 15.114v8.802h5v-8.802l-.833-1.761z"/><path fill="#d22f27" d="M40.998 14.283l-3.57-.061l-3.571-.061l1.838-3.08L37.532 8l1.733 3.141l1.733 3.142z"/><path fill="#9b9b9a" d="M26 22.86h10.5V63H26z"/><path fill="#d0cfce" d="M36.5 22.86H47V63H36.5z"/><path fill="#9b9b9a" d="M31.25 18H28.1l-1.05 1.2V24h4.2v-2.4l1.05-2.4l-1.05-1.2z"/><path fill="#d22f27" d="M33.484 17.449l-3.489-.06l-3.488-.059l1.795-3.01l1.796-3.009l1.693 3.069l1.693 3.069z"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M46.949 18.484l-2.871-.049l-2.871-.049l1.478-2.513l1.478-2.514l1.393 2.562l1.393 2.563z"/><path stroke-linecap="round" d="M40.7 19.691v-4.489l-1.05-1.849h-4.2l-1.05 1.849v4.489"/><path stroke-linecap="round" d="M40.998 14.283l-3.57-.061l-3.571-.061l1.838-3.08L37.532 8l1.733 3.141l1.733 3.142z"/><path stroke-linecap="round" d="M26 22.86h10.5V63H26z"/><path stroke-linecap="round" d="M36.5 22.86H47V63H36.5z"/><path stroke-linecap="round" d="M39.65 27.085v2.113"/><path stroke-linecap="round" d="M42.8 37.648v2.113"/><path stroke-linecap="round" d="M39.65 50.324v2.113"/><path stroke-linecap="round" d="M27.05 17.579v4.225"/><path d="M33.484 17.449l-3.489-.06l-3.488-.059l1.795-3.01l1.796-3.009l1.693 3.069l1.693 3.069z"/><path stroke-linecap="round" d="M45.95 22.86v-4.225"/><path stroke-linecap="round" d="M30.2 37.648v2.113"/><path stroke-linecap="round" d="M33.35 50.324v2.113"/><path stroke-linecap="round" d="M37.55 17.579h.01"/></g>`),
		g.Group(children),
	)
}