package core

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"image"
	"os"
	"path/filepath"

	nn_util "github.com/saent-x/ids-nn/core"
	"github.com/saent-x/ids-nn/core/model"
)

type InferenceResults struct {
	ModelPrediction string
}

func RunInference(img image.Image) (InferenceResults, error) {
	img = nn_util.ConvertIntoGrayscale(img, 28, 28)

	imgBytes := img.(*image.Gray)

	imgSlice, err := nn_util.ScaleValues(imgBytes, true)
	if err != nil {
		return InferenceResults{}, err
	}

	image_data := nn_util.CreateDenseMatrix(1, len(imgSlice), imgSlice)

	savedModels, err1 := filepath.Abs("core/saved_models")
	if err1 != nil {
		return InferenceResults{}, err1
	}

	modelFile, err := os.Open(fmt.Sprintf("%v/%v.json", savedModels, "fashionMNIST_model_full"))
	if err != nil {
		panic(err)
	}
	defer modelFile.Close()

	modelDataProvider := new(model.ModelDataProvider)
	savedModel, err := modelDataProvider.Load(modelFile)
	if err != nil {
		return InferenceResults{}, err
	}

	confidences := savedModel.Predict(image_data, 0)
	predictions := savedModel.OutputLayerActivation.Predictions(confidences)

	fmt.Println(mat.Formatted(predictions))

	fashionMNIST_labels := map[int]string{
		0: "T-shirt/top",
		1: "Trouser",
		2: "Pullover",
		3: "Dress",
		4: "Coat",
		5: "Sandal",
		6: "Shirt",
		7: "Sneaker",
		8: "Bag",
		9: "Ankle boot",
	}

	return InferenceResults{
		ModelPrediction: fashionMNIST_labels[int(predictions.At(0, 0))],
	}, nil
}
