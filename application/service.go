package application

import (
	context "context"
	"google.golang.org/grpc"
	"github.com/hcliff-zhang/playground/server/serverpb"
)

// Service wraps a gRPC client for the Api service and provides methods to interact with the database via RPC.
type Service struct {
	client serverpb.ApiClient
}

func NewService(conn *grpc.ClientConn) *Service {
	return &Service{client: serverpb.NewApiClient(conn)}
}

// CreatePatient sends a CreatePatientRequest via gRPC and returns the created patient.
func (s *Service) CreatePatient(ctx context.Context, patient *serverpb.Patient) (*serverpb.Patient, error) {
	resp, err := s.client.CreatePatient(ctx, &serverpb.CreatePatientRequest{Patient: patient})
	if err != nil {
		return nil, err
	}
	return resp.Patient, nil
}

// GetPatient fetches a patient by ID via gRPC.
func (s *Service) GetPatient(ctx context.Context, id uint64) (*serverpb.Patient, error) {
	resp, err := s.client.GetPatient(ctx, &serverpb.GetPatientRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return resp.Patient, nil
}

// ListPatients fetches a paginated list of patients via gRPC.
func (s *Service) ListPatients(ctx context.Context, limit, offset int32) ([]*serverpb.Patient, error) {
	resp, err := s.client.ListPatients(ctx, &serverpb.ListPatientsRequest{Limit: limit, Offset: offset})
	if err != nil {
		return nil, err
	}
	return resp.Patients, nil
}

// CreatePrescription sends a CreatePrescriptionRequest via gRPC and returns the created prescription.
func (s *Service) CreatePrescription(ctx context.Context, patientID uint64, prescription *serverpb.Prescription) (*serverpb.Prescription, error) {
	resp, err := s.client.CreatePrescription(ctx, &serverpb.CreatePrescriptionRequest{PatientId: patientID, Prescription: prescription})
	if err != nil {
		return nil, err
	}
	return resp.Prescription, nil
}

// GetPrescription fetches a prescription by ID via gRPC.
func (s *Service) GetPrescription(ctx context.Context, id uint64) (*serverpb.Prescription, error) {
	resp, err := s.client.GetPrescription(ctx, &serverpb.GetPrescriptionRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return resp.Prescription, nil
}

// ListPrescriptionsForPatient fetches all prescriptions for a patient via gRPC.
func (s *Service) ListPrescriptionsForPatient(ctx context.Context, patientID uint64) ([]*serverpb.Prescription, error) {
	resp, err := s.client.ListPrescriptionsForPatient(ctx, &serverpb.ListPrescriptionsForPatientRequest{PatientId: patientID})
	if err != nil {
		return nil, err
	}
	return resp.Prescriptions, nil
}
