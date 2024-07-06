package app

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/vrv501/go-repo/grpc/generated/bank"
	pbdate "github.com/vrv501/go-repo/grpc/generated/date"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type bankService struct {
	pb.BankServiceServer
}

func NewBankService() *bankService {
	return &bankService{}
}

func (bs *bankService) GetCurrentBalance(ctx context.Context, in *pb.CurrentBalanceRequest) (*pb.CurrentBalanceResponse, error) {
	log.Printf("invoked GetCurrentBalance for %v\n", in.AccountNumber)

	if in.AccountNumber == "0987" {
		return nil, status.Error(codes.NotFound, "Adasdasda")
	}

	return &pb.CurrentBalanceResponse{
		Amount: 40,
		CurrentDate: &pbdate.Date{
			Year:  2022,
			Month: 12,
			Day:   3,
		},
	}, nil
}

func (bs *bankService) FetchExchangeRates(in *pb.ExchangeRateRequest, stream pb.BankService_FetchExchangeRatesServer) error {
	log.Printf("invoked FetchExchangeRates for from:%v to: %v\n", in.FromCurr, in.ToCurr)

	tme, ok := stream.Context().Deadline()
	fmt.Println(time.Now())
	fmt.Println(time.Now().Add(180 * time.Second))
	fmt.Println(tme, ok)

	ctx := stream.Context()
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Println(md.Get("authorization")[0])
	}

	err := stream.SendHeader(metadata.New(map[string]string{"ASdasd": "asdasd"}))
	if err != nil {
		return err
	}

	for i := range 3 {
		err := stream.Send(&pb.ExchangeRateResponse{
			Request: in,
			Rate:    float64(i),
			Timestamp: &timestamppb.Timestamp{
				Seconds: timestamppb.Now().Seconds,
				Nanos:   timestamppb.Now().Nanos,
			},
		})
		if err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}

func (bs *bankService) SummarizeTransactions(stream pb.BankService_SummarizeTransactionsServer) error {

	initAmount := 2000
	inAmount := 0
	outAmount := 0
	accNumb := ""
	for {

		req, err := stream.Recv()
		if err != nil {
			log.Println("invoked SummarizeTransactions stopped")
			if err == io.EOF {
				break
			}
			return err
		}
		log.Printf("invoked SummarizeTransactions at: %v\n", req.Timestamp)
		accNumb = req.AccountNumber
		if req.Type == pb.TransactionType_TRANSACTION_TYPE_OUT {
			outAmount += int(req.Amount)
		} else if req.Type == pb.TransactionType_TRANSACTION_TYPE_IN {
			inAmount += int(req.Amount)
		}
	}

	err := stream.SendAndClose(&pb.TransactionSummary{
		AccountNumber: accNumb,
		TotalMoneyIn:  float64(inAmount),
		TotalMoneyOut: float64(outAmount),
		Balance:       float64(initAmount) + float64(inAmount) - float64(outAmount),
		TransactionDate: &pbdate.Date{
			Year:  2022,
			Month: 20,
			Day:   02,
		},
	})
	if err != nil {
		return err
	}
	return nil
}
