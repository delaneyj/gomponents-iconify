package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Folder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaFolder0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-miterlimit="4" stroke-width="16"><path id="galaFolder1" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="M 240.04121,64.000608 240,208 c -0.005,15.99988 -16.04134,32 -32,32 H 48 C 31.958138,240 15.996182,223.99988 16,208 l 0.04199,-175.999258"/><path id="galaFolder2" stroke-linecap="round" stroke-linejoin="round" d="m 224,48 c 8.83652,8e-6 15.99999,7.163485 16,16"/><path id="galaFolder3" stroke-linecap="butt" stroke-linejoin="miter" stroke-opacity="1" d="m 144.04154,48.000675 h 79.99974"/><path id="galaFolder4" stroke-linecap="round" stroke-linejoin="round" d="m 106.38483,16.000772 a 15.999945,15.999945 0 0 1 13.85635,7.999985"/><path id="galaFolder5" stroke-linecap="round" stroke-linejoin="round" d="M 144.04154,48.000675 A 15.999945,15.999945 0 0 1 130.18519,40.00069"/><path id="galaFolder6" stroke-linecap="round" stroke-linejoin="round" d="M 144.04154,48.000675 A 15.999945,15.999945 0 0 0 130.18519,56.00066"/><path id="galaFolder7" stroke-linecap="butt" stroke-linejoin="miter" stroke-opacity="1" d="m 120.24118,24.000757 9.94401,15.999933"/><path id="galaFolder8" stroke-linecap="round" stroke-linejoin="round" d="m 106.3848,80.000579 a 15.999945,15.999945 0 0 0 13.85638,-7.999986"/><path id="galaFolder9" stroke-linecap="butt" stroke-linejoin="miter" stroke-opacity="1" d="M 120.24118,72.000593 130.18519,56.00066"/><path id="galaFoldera" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="m 16,80 90.38472,5.79e-4"/><path id="galaFolderb" stroke-linecap="round" stroke-linejoin="round" d="m 32.041967,16.000772 c -8.836533,3e-6 -15.999992,7.167628 -15.999978,16.004161"/><path id="galaFolderc" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="M 32.041967,16.000772 H 106.38498"/></g>`),
		g.Group(children),
	)
}