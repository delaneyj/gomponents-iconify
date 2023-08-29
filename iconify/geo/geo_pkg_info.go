package geo

import (
	g "github.com/maragudk/gomponents"
)

const IconifyVersion = "0.0.10"

func IconFromName(name string) g.Node {
	switch name {
	case "turfAlong": return TurfAlong()
	case "turfBboxPolygon": return TurfBboxPolygon()
	case "turfBezier": return TurfBezier()
	case "turfBuffer": return TurfBuffer()
	case "turfCenter": return TurfCenter()
	case "turfCentroid": return TurfCentroid()
	case "turfConcave": return TurfConcave()
	case "turfConvex": return TurfConvex()
	case "turfDestination": return TurfDestination()
	case "turfEnvelope": return TurfEnvelope()
	case "turfErased": return TurfErased()
	case "turfExplode": return TurfExplode()
	case "turfExtent": return TurfExtent()
	case "turfIntersect": return TurfIntersect()
	case "turfKinks": return TurfKinks()
	case "turfLineSlice": return TurfLineSlice()
	case "turfMerge": return TurfMerge()
	case "turfMidpoint": return TurfMidpoint()
	case "turfPointGrid": return TurfPointGrid()
	case "turfPointOnLine": return TurfPointOnLine()
	case "turfPointOnSurface": return TurfPointOnSurface()
	case "turfSimplify": return TurfSimplify()
	case "turfSize": return TurfSize()
	case "turfSquare": return TurfSquare()
	case "turfSquareGrid": return TurfSquareGrid()
	case "turfTin": return TurfTin()
	case "turfTriangleGrid": return TurfTriangleGrid()
	case "turfUnion": return TurfUnion()
	case "uiEarthEast": return UiEarthEast()
	case "uiEarthWest": return UiEarthWest()
	default:
		panic("Unknown icon name: " + name)
	}
}
