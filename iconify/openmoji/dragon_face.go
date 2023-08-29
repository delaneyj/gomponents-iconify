package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DragonFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#5C9E31" stroke="#5C9E31" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="1.8" d="M51 33c8-4 8-9 8-9l-7 3c4-5 2-11 2-11c-2 4-5 5-5 5c-4-6-10-6-10-6l-3-4l-3 4c-6 0-10 6-10 6c-4-1-5-5-5-5c-1 9 3 11 3 11l-7-2c2 4 7 7 7 7c-3 2-6 1-6 1c3 5 7 5 7 5h28c7 0 8-5 8-5c-5 1-7 0-7 0z"/><path fill="#B1CC33" stroke="#B1CC33" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M26 38s0-1 4.054-.929c0 0 5.947-3.964 11.893 0c0 0 4.053.947 4.053 2.929c0 0 4 3 1.894 5.992c0 0 0 3.964-5.946 7.929c0 0-3.965 1.982-11.894 0c0 0-7.267-6.07-6.054-10.921c1-4 2-5 2-5z"/><path fill="#EA5A47" stroke="#EA5A47" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M34 49s0 10-2 11c0 0 4-1 4-5c0 0 1 4 4 5c0 0-3-7-2-11h-4z"/><path fill="#FFF" d="M31 46v4c-3-3-3-7-3-7s8 7 16 0c0 0 1 5-3 7v-4"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M31 46v4c-3-3-3-7-3-7s8 7 16 0c0 0 1 5-3 7v-4M27 23s7 2 9 8c0 0 2-7 10-8"/><circle cx="29" cy="30" r="2"/><circle cx="43" cy="30" r="2"/><circle cx="34" cy="39" r="1"/><circle cx="38" cy="39" r="1"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M34 49s0 10-2 11c0 0 4-1 4-5c0 0 1 4 4 5c0 0-3-7-2-11M22 38s-4 0-7-5c0 0 3 1 6-1c0 0-5-3-7-7l7 2s-4-2-3-11c0 0 1 4 5 5c0 0 4-6 10-6l3-4l3 4s6 0 10 6c0 0 3-1 5-5c0 0 2 6-2 11l7-3s0 5-8 9c0 0 2 1 7 0c0 0-1 5-8 5"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M25 39s-1-3 5-2c0 0 6-5 12 0c0 0 5-1 5 2m-18 1s-6-1-9 3s-6 4-8 2m31-5s6-1 9 3s6 4 8 2"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M23.315 40.571S22 51 30 54m19.189-13.217S49 51 42 54m-11-8v4c-3-3-3-7-3-7s8 7 16 0c0 0 1 5-3 7v-4"/>`),
		g.Group(children),
	)
}