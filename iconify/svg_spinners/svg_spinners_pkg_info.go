package svg_spinners

import (
	g "github.com/maragudk/gomponents"
)

const IconifyVersion = ""

func IconFromName(name string) g.Node {
	switch name {
	case "barsFade": return BarsFade()
	case "barsRotateFade": return BarsRotateFade()
	case "barsScale": return BarsScale()
	case "barsScaleFade": return BarsScaleFade()
	case "barsScaleMiddle": return BarsScaleMiddle()
	case "blocksScale": return BlocksScale()
	case "blocksShuffleThree": return BlocksShuffleThree()
	case "blocksShuffleTwo": return BlocksShuffleTwo()
	case "blocksWave": return BlocksWave()
	case "bouncingBall": return BouncingBall()
	case "clock": return Clock()
	case "dotRevolve": return DotRevolve()
	case "eclipse": return Eclipse()
	case "eclipseHalf": return EclipseHalf()
	case "eightDotsRotate": return EightDotsRotate()
	case "gooeyBallsOne": return GooeyBallsOne()
	case "gooeyBallsTwo": return GooeyBallsTwo()
	case "ninetyRing": return NinetyRing()
	case "ninetyRingWithBg": return NinetyRingWithBg()
	case "oneHundredEightyRing": return OneHundredEightyRing()
	case "oneHundredEightyRingWithBg": return OneHundredEightyRingWithBg()
	case "pulse": return Pulse()
	case "pulseMultiple": return PulseMultiple()
	case "pulseRing": return PulseRing()
	case "pulseRingsMultiple": return PulseRingsMultiple()
	case "pulseRingsThree": return PulseRingsThree()
	case "pulseRingsTwo": return PulseRingsTwo()
	case "pulseThree": return PulseThree()
	case "pulseTwo": return PulseTwo()
	case "ringResize": return RingResize()
	case "sixDotsRotate": return SixDotsRotate()
	case "sixDotsScale": return SixDotsScale()
	case "sixDotsScaleMiddle": return SixDotsScaleMiddle()
	case "tadpole": return Tadpole()
	case "threeDotsBounce": return ThreeDotsBounce()
	case "threeDotsFade": return ThreeDotsFade()
	case "threeDotsMove": return ThreeDotsMove()
	case "threeDotsRotate": return ThreeDotsRotate()
	case "threeDotsScale": return ThreeDotsScale()
	case "threeDotsScaleMiddle": return ThreeDotsScaleMiddle()
	case "twelveDotsScaleRotate": return TwelveDotsScaleRotate()
	case "twoHundredSeventyRing": return TwoHundredSeventyRing()
	case "twoHundredSeventyRingWithBg": return TwoHundredSeventyRingWithBg()
	case "wifi": return Wifi()
	case "wifiFade": return WifiFade()
	case "windToy": return WindToy()
	default:
		panic("Unknown icon name: " + name)
	}
}
