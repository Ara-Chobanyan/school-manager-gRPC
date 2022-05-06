package api

import (
	"context"
	"log"

	pb "github.com/Ara-Chobanyan/school-manager/global-grade-service/api/proto"
)

const red = "RED"
const orange = "Orange"
const green = "Green"

func (r *Adapter) GetGlobalAverage(ctx context.Context, in *pb.GlobalParams) (*pb.GlobalAverage, error) {

	average, err := r.api.GetGlobalAverage()

	if err != nil {
		log.Printf("Could not get status: %v", err)
		return nil, err
	}

	res := &pb.GlobalAverage{
		Grade:    float32(average.Grade),
		Behavior: float32(average.Behaviour),
	}

	switch {
	case res.Grade < 50:
		res.Status = red
	case res.Grade <= 70:
		res.Status = orange
	case res.Grade >= 70:
		res.Status = green
	}

	return res, nil
}
