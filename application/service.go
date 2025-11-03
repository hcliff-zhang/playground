package application

import (
	"context"

	"github.com/hcliff-zhang/playground/database"
	"github.com/hcliff-zhang/playground/server/serverpb"
)

// Service wraps a database handle and provides methods to read and write data.
type Service struct {
	serverpb.UnimplementedApiServer
	DB *database.DB
}

func NewService(db *database.DB) *Service {
	return &Service{DB: db}
}

// --- Patient methods ---

// CreatePatient creates a new patient in the database.
func (s *Service) CreatePatient(ctx context.Context, req *serverpb.CreatePatientRequest) (*serverpb.CreatePatientResponse, error) {
	// Convert proto to database model
	dbPatient := PatientFromProto(req.Patient)
	
	// Save to database
	if err := s.DB.CreatePatient(dbPatient); err != nil {
		return nil, err
	}
	
	// Convert back to proto
	return &serverpb.CreatePatientResponse{
		Patient: PatientToProto(dbPatient),
	}, nil
}

// GetPatient fetches a patient by ID with preloaded prescriptions.
func (s *Service) GetPatient(ctx context.Context, req *serverpb.GetPatientRequest) (*serverpb.GetPatientResponse, error) {
	dbPatient, err := s.DB.GetPatientByID(uint(req.Id))
	if err != nil {
		return nil, err
	}
	
	return &serverpb.GetPatientResponse{
		Patient: PatientToProto(dbPatient),
	}, nil
}

// ListPatients returns a paginated list of patients.
func (s *Service) ListPatients(ctx context.Context, req *serverpb.ListPatientsRequest) (*serverpb.ListPatientsResponse, error) {
	dbPatients, err := s.DB.ListPatients(int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}
	
	return &serverpb.ListPatientsResponse{
		Patients: PatientsToProto(dbPatients),
		Total:    int32(len(dbPatients)),
	}, nil
}

// --- Prescription methods ---

// CreatePrescription creates a prescription associated with a patient.
func (s *Service) CreatePrescription(ctx context.Context, req *serverpb.CreatePrescriptionRequest) (*serverpb.CreatePrescriptionResponse, error) {
	// Convert proto to database model
	dbPrescription := PrescriptionFromProto(req.Prescription)
	
	// Save to database
	if err := s.DB.CreatePrescriptionForPatient(uint(req.PatientId), dbPrescription); err != nil {
		return nil, err
	}
	
	// Convert back to proto
	return &serverpb.CreatePrescriptionResponse{
		Prescription: PrescriptionToProto(dbPrescription),
	}, nil
}

// GetPrescription fetches a prescription by ID.
func (s *Service) GetPrescription(ctx context.Context, req *serverpb.GetPrescriptionRequest) (*serverpb.GetPrescriptionResponse, error) {
	dbPrescription, err := s.DB.GetPrescriptionByID(uint(req.Id))
	if err != nil {
		return nil, err
	}
	
	return &serverpb.GetPrescriptionResponse{
		Prescription: PrescriptionToProto(dbPrescription),
	}, nil
}

// ListPrescriptionsForPatient returns all prescriptions for a patient.
func (s *Service) ListPrescriptionsForPatient(ctx context.Context, req *serverpb.ListPrescriptionsForPatientRequest) (*serverpb.ListPrescriptionsResponse, error) {
	dbPrescriptions, err := s.DB.ListPrescriptionsForPatientAssoc(uint(req.PatientId))
	if err != nil {
		return nil, err
	}
	
	return &serverpb.ListPrescriptionsResponse{
		Prescriptions: PrescriptionsToProto(dbPrescriptions),
	}, nil
}
