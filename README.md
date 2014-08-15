```bash
$ go test -bench . -cpu 1,2,4 -timeout=30m
```

```text
BenchmarkScalerNfntResizeNearestNeighbor-2	       1	1143478510 ns/op
BenchmarkScalerNfntResizeNearestNeighbor-4	       2	 901209225 ns/op
BenchmarkScalerNfntResizeBilinear	       1	2150473395 ns/op
BenchmarkScalerNfntResizeBilinear-2	       1	1532324746 ns/op
BenchmarkScalerNfntResizeBilinear-4	       1	1030419092 ns/op
BenchmarkScalerNfntResizeBicubic	       1	3519844528 ns/op
BenchmarkScalerNfntResizeBicubic-2	       1	1996684355 ns/op
BenchmarkScalerNfntResizeBicubic-4	       1	1439740487 ns/op
BenchmarkScalerNfntResizeMitchellNetravali	       1	3488235550 ns/op
BenchmarkScalerNfntResizeMitchellNetravali-2	       1	1981832815 ns/op
BenchmarkScalerNfntResizeMitchellNetravali-4	       1	1453293607 ns/op
BenchmarkScalerNfntResizeLanczos2	       1	3500707388 ns/op
BenchmarkScalerNfntResizeLanczos2-2	       1	2002036013 ns/op
BenchmarkScalerNfntResizeLanczos2-4	       1	1452755828 ns/op
BenchmarkScalerNfntResizeLanczos3	       1	4886655419 ns/op
BenchmarkScalerNfntResizeLanczos3-2	       1	2708343359 ns/op
BenchmarkScalerNfntResizeLanczos3-4	       1	1915247532 ns/op
BenchmarkScalerGiftResizeNearestNeighbor	      50	  49171325 ns/op
BenchmarkScalerGiftResizeNearestNeighbor-2	     100	  24908094 ns/op
BenchmarkScalerGiftResizeNearestNeighbor-4	     100	  16480055 ns/op
BenchmarkScalerGiftResizeLinear	       1	6000873715 ns/op
BenchmarkScalerGiftResizeLinear-2	       1	3052068772 ns/op
BenchmarkScalerGiftResizeLinear-4	       1	2009008547 ns/op
BenchmarkScalerGiftResizeBox	       1	5240528492 ns/op
BenchmarkScalerGiftResizeBox-2	       1	2665748067 ns/op
BenchmarkScalerGiftResizeBox-4	       1	1860051687 ns/op
BenchmarkScalerGiftResizeCubic	       1	7456741810 ns/op
BenchmarkScalerGiftResizeCubic-2	       1	3820033889 ns/op
BenchmarkScalerGiftResizeCubic-4	       1	2458354376 ns/op
BenchmarkScalerGiftResizeLanczos	       1	8995416247 ns/op
BenchmarkScalerGiftResizeLanczos-2	       1	4572013802 ns/op
BenchmarkScalerGiftResizeLanczos-4	       1	2893705687 ns/op
BenchmarkScalerMoustachioResample	      50	  59382551 ns/op
BenchmarkScalerMoustachioResample-2	      50	  59577875 ns/op
BenchmarkScalerMoustachioResample-4	      50	  59442952 ns/op
BenchmarkScalerMoustachioResize	       1	8728855744 ns/op
BenchmarkScalerMoustachioResize-2	       1	8719529697 ns/op
BenchmarkScalerMoustachioResize-4	       1	8713259278 ns/op
BenchmarkScalerRezConvertBilinear	      20	 124310078 ns/op
BenchmarkScalerRezConvertBilinear-2	      50	  66188208 ns/op
BenchmarkScalerRezConvertBilinear-4	      50	  50397984 ns/op
BenchmarkScalerRezConvertBicubic	       5	 262255626 ns/op
BenchmarkScalerRezConvertBicubic-2	      10	 140172601 ns/op
BenchmarkScalerRezConvertBicubic-4	      20	 114420912 ns/op
BenchmarkScalerRezConvertLanczos2	       5	 291210914 ns/op
BenchmarkScalerRezConvertLanczos2-2	      10	 158286294 ns/op
BenchmarkScalerRezConvertLanczos2-4	      20	 123242892 ns/op
BenchmarkScalerRezConvertLanczos3	       5	 491630269 ns/op
BenchmarkScalerRezConvertLanczos3-2	       5	 260901948 ns/op
BenchmarkScalerRezConvertLanczos3-4	      10	 205677418 ns/op
BenchmarkScalerGraphicsScale	       5	 270507122 ns/op
BenchmarkScalerGraphicsScale-2	       5	 270752502 ns/op
BenchmarkScalerGraphicsScale-4	       5	 270122618 ns/op
BenchmarkScalerImagingResizeNearestNeighbor	       1	2601204874 ns/op
BenchmarkScalerImagingResizeNearestNeighbor-2	       1	2535275541 ns/op
BenchmarkScalerImagingResizeNearestNeighbor-4	       1	2541598411 ns/op
BenchmarkScalerImagingResizeBox	       1	4869684336 ns/op
BenchmarkScalerImagingResizeBox-2	       1	3722431114 ns/op
BenchmarkScalerImagingResizeBox-4	       1	3198125893 ns/op
BenchmarkScalerImagingResizeLinear	       1	6492370852 ns/op
BenchmarkScalerImagingResizeLinear-2	       1	4541014608 ns/op
BenchmarkScalerImagingResizeLinear-4	       1	3629865432 ns/op
BenchmarkScalerImagingResizeHermite	       1	6495070078 ns/op
BenchmarkScalerImagingResizeHermite-2	       1	4536131585 ns/op
BenchmarkScalerImagingResizeHermite-4	       1	3666585185 ns/op
BenchmarkScalerImagingResizeMitchellNetravali	       1	9551198515 ns/op
BenchmarkScalerImagingResizeMitchellNetravali-2	       1	6067452432 ns/op
BenchmarkScalerImagingResizeMitchellNetravali-4	       1	4541493301 ns/op
BenchmarkScalerImagingResizeCatmullRom	       1	9538244075 ns/op
BenchmarkScalerImagingResizeCatmullRom-2	       1	6087169605 ns/op
BenchmarkScalerImagingResizeCatmullRom-4	       1	4538231154 ns/op
BenchmarkScalerImagingResizeBSpline	       1	9549585042 ns/op
BenchmarkScalerImagingResizeBSpline-2	       1	6072112522 ns/op
BenchmarkScalerImagingResizeBSpline-4	       1	4538819496 ns/op
BenchmarkScalerImagingResizeGaussian	       1	9547487568 ns/op
BenchmarkScalerImagingResizeGaussian-2	       1	6075915800 ns/op
BenchmarkScalerImagingResizeGaussian-4	       1	4549634978 ns/op
BenchmarkScalerImagingResizeBartlett	       1	12763536202 ns/op
BenchmarkScalerImagingResizeBartlett-2	       1	7698019281 ns/op
BenchmarkScalerImagingResizeBartlett-4	       1	5488685119 ns/op
BenchmarkScalerImagingResizeLanczos	       1	12756529026 ns/op
BenchmarkScalerImagingResizeLanczos-2	       1	7684333881 ns/op
BenchmarkScalerImagingResizeLanczos-4	       1	5489842598 ns/op
BenchmarkScalerImagingResizeHann	       1	12769159193 ns/op
BenchmarkScalerImagingResizeHann-2	       1	7701759421 ns/op
BenchmarkScalerImagingResizeHann-4	       1	5498424626 ns/op
BenchmarkScalerImagingResizeHamming	       1	12753637527 ns/op
BenchmarkScalerImagingResizeHamming-2	       1	7691755425 ns/op
BenchmarkScalerImagingResizeHamming-4	       1	5505454360 ns/op
BenchmarkScalerImagingResizeWelch	       1	12750364484 ns/op
BenchmarkScalerImagingResizeWelch-2	       1	7688277259 ns/op
BenchmarkScalerImagingResizeWelch-4	       1	5491082281 ns/op
BenchmarkScalerImagingResizeCosine	       1	12764634316 ns/op
BenchmarkScalerImagingResizeCosine-2	       1	7693310162 ns/op
BenchmarkScalerImagingResizeCosine-4	       1	5497780409 ns/op
BenchmarkFlowNfntResizeNearestNeighbor	       1	6124568424 ns/op
BenchmarkFlowNfntResizeNearestNeighbor-2	       1	5434221419 ns/op
BenchmarkFlowNfntResizeNearestNeighbor-4	       1	5148405604 ns/op
BenchmarkFlowNfntResizeBilinear	       1	6428653671 ns/op
BenchmarkFlowNfntResizeBilinear-2	       1	5607631948 ns/op
BenchmarkFlowNfntResizeBilinear-4	       1	5286887111 ns/op
BenchmarkFlowNfntResizeBicubic	       1	7816138212 ns/op
BenchmarkFlowNfntResizeBicubic-2	       1	6278170360 ns/op
BenchmarkFlowNfntResizeBicubic-4	       1	5747701925 ns/op
BenchmarkFlowNfntResizeMitchellNetravali	       1	7755977069 ns/op
BenchmarkFlowNfntResizeMitchellNetravali-2	       1	6261916128 ns/op
BenchmarkFlowNfntResizeMitchellNetravali-4	       1	5719189641 ns/op
BenchmarkFlowNfntResizeLanczos2	       1	7782603501 ns/op
BenchmarkFlowNfntResizeLanczos2-2	       1	6280996138 ns/op
BenchmarkFlowNfntResizeLanczos2-4	       1	5731810100 ns/op
BenchmarkFlowNfntResizeLanczos3	       1	9170899526 ns/op
BenchmarkFlowNfntResizeLanczos3-2	       1	6979949895 ns/op
BenchmarkFlowNfntResizeLanczos3-4	       1	6182520696 ns/op
BenchmarkFlowGiftResizeNearestNeighbor	       1	3954503572 ns/op
BenchmarkFlowGiftResizeNearestNeighbor-2	       1	3936756622 ns/op
BenchmarkFlowGiftResizeNearestNeighbor-4	       1	3916828181 ns/op
BenchmarkFlowGiftResizeLinear	       1	9983205826 ns/op
BenchmarkFlowGiftResizeLinear-2	       1	7065667290 ns/op
BenchmarkFlowGiftResizeLinear-4	       1	6045515565 ns/op
BenchmarkFlowGiftResizeBox	       1	9231263906 ns/op
BenchmarkFlowGiftResizeBox-2	       1	6647171962 ns/op
BenchmarkFlowGiftResizeBox-4	       1	5787244942 ns/op
BenchmarkFlowGiftResizeCubic	       1	11431261406 ns/op
BenchmarkFlowGiftResizeCubic-2	       1	7793885382 ns/op
BenchmarkFlowGiftResizeCubic-4	       1	6460391088 ns/op
BenchmarkFlowGiftResizeLanczos	       1	12963771232 ns/op
BenchmarkFlowGiftResizeLanczos-2	       1	8561452682 ns/op
BenchmarkFlowGiftResizeLanczos-4	       1	6861766779 ns/op
BenchmarkFlowMoustachioResample	       1	3964251670 ns/op
BenchmarkFlowMoustachioResample-2	       1	3970976058 ns/op
BenchmarkFlowMoustachioResample-4	       1	3966966819 ns/op
BenchmarkFlowMoustachioResize	       1	12684947786 ns/op
BenchmarkFlowMoustachioResize-2	       1	12693481343 ns/op
BenchmarkFlowMoustachioResize-4	       1	12695206423 ns/op
BenchmarkFlowRezConvertBilinear	       1	4380359542 ns/op
BenchmarkFlowRezConvertBilinear-2	       1	4345087808 ns/op
BenchmarkFlowRezConvertBilinear-4	       1	4305416352 ns/op
BenchmarkFlowRezConvertBicubic	       1	4508385039 ns/op
BenchmarkFlowRezConvertBicubic-2	       1	4458174395 ns/op
BenchmarkFlowRezConvertBicubic-4	       1	4375709719 ns/op
BenchmarkFlowRezConvertLanczos2	       1	4558291736 ns/op
BenchmarkFlowRezConvertLanczos2-2	       1	4431371266 ns/op
BenchmarkFlowRezConvertLanczos2-4	       1	4395763386 ns/op
BenchmarkFlowRezConvertLanczos3	       1	4735674138 ns/op
BenchmarkFlowRezConvertLanczos3-2	       1	4520838858 ns/op
BenchmarkFlowRezConvertLanczos3-4	       1	4483394991 ns/op
BenchmarkFlowGraphicsScale	       1	4179281248 ns/op
BenchmarkFlowGraphicsScale-2	       1	4185824329 ns/op
BenchmarkFlowGraphicsScale-4	       1	4183316607 ns/op
BenchmarkFlowImagingResizeNearestNeighbor	       1	6450873397 ns/op
BenchmarkFlowImagingResizeNearestNeighbor-2	       1	6450371867 ns/op
BenchmarkFlowImagingResizeNearestNeighbor-4	       1	6442784938 ns/op
BenchmarkFlowImagingResizeBox	       1	8852432981 ns/op
BenchmarkFlowImagingResizeBox-2	       1	7712104941 ns/op
BenchmarkFlowImagingResizeBox-4	       1	7200776300 ns/op
BenchmarkFlowImagingResizeLinear	       1	10501894200 ns/op
BenchmarkFlowImagingResizeLinear-2	       1	8561606961 ns/op
BenchmarkFlowImagingResizeLinear-4	       1	7693973305 ns/op
BenchmarkFlowImagingResizeHermite	       1	10494328365 ns/op
BenchmarkFlowImagingResizeHermite-2	       1	8548951372 ns/op
BenchmarkFlowImagingResizeHermite-4	       1	7654592809 ns/op
BenchmarkFlowImagingResizeMitchellNetravali	       1	13552661062 ns/op
BenchmarkFlowImagingResizeMitchellNetravali-2	       1	10096533855 ns/op
BenchmarkFlowImagingResizeMitchellNetravali-4	       1	8543969383 ns/op
BenchmarkFlowImagingResizeCatmullRom	       1	13528374453 ns/op
BenchmarkFlowImagingResizeCatmullRom-2	       1	10083751014 ns/op
BenchmarkFlowImagingResizeCatmullRom-4	       1	8522643766 ns/op
BenchmarkFlowImagingResizeBSpline	       1	13541974228 ns/op
BenchmarkFlowImagingResizeBSpline-2	       1	10100896001 ns/op
BenchmarkFlowImagingResizeBSpline-4	       1	8614686267 ns/op
BenchmarkFlowImagingResizeGaussian	       1	13561140492 ns/op
BenchmarkFlowImagingResizeGaussian-2	       1	10094559791 ns/op
BenchmarkFlowImagingResizeGaussian-4	       1	8592674040 ns/op
BenchmarkFlowImagingResizeBartlett	       1	16731082144 ns/op
BenchmarkFlowImagingResizeBartlett-2	       1	11667487534 ns/op
BenchmarkFlowImagingResizeBartlett-4	       1	9462249704 ns/op
BenchmarkFlowImagingResizeLanczos	       1	16724123299 ns/op
BenchmarkFlowImagingResizeLanczos-2	       1	11666379363 ns/op
BenchmarkFlowImagingResizeLanczos-4	       1	9466969578 ns/op
BenchmarkFlowImagingResizeHann	       1	16751070300 ns/op
BenchmarkFlowImagingResizeHann-2	       1	11708008261 ns/op
BenchmarkFlowImagingResizeHann-4	       1	9469059159 ns/op
BenchmarkFlowImagingResizeHamming	       1	16751608777 ns/op
BenchmarkFlowImagingResizeHamming-2	       1	11692684873 ns/op
BenchmarkFlowImagingResizeHamming-4	       1	9461217380 ns/op
BenchmarkFlowImagingResizeWelch	       1	16726808836 ns/op
BenchmarkFlowImagingResizeWelch-2	       1	11671770621 ns/op
BenchmarkFlowImagingResizeWelch-4	       1	9444504640 ns/op
BenchmarkFlowImagingResizeCosine	       1	16703786354 ns/op
BenchmarkFlowImagingResizeCosine-2	       1	11661413648 ns/op
BenchmarkFlowImagingResizeCosine-4	       1	9463419127 ns/op
BenchmarkFlowImageMagickResize	       1	3294664357 ns/op
BenchmarkFlowImageMagickResize-2	       1	3267215265 ns/op
BenchmarkFlowImageMagickResize-4	       1	3295412287 ns/op
BenchmarkFlowImageMagickResizeWithDefine	       5	 482721371 ns/op
BenchmarkFlowImageMagickResizeWithDefine-2	       5	 481486116 ns/op
BenchmarkFlowImageMagickResizeWithDefine-4	       5	 484040840 ns/op
ok  	github.com/oov/resize-test	1269.408s
```
