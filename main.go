package main

import (
	"fmt"
	"go-jobs/models"
	"go-jobs/store"
	"time"
)

func main() {
	s := store.New()

	fmt.Println("--- Testing Job System Store ---")

	// TEST
	jobID := "job-123"
	job1 := &models.Job{
		ID:         jobID,
		Status:     models.StatusPending,
		Payload:    "Send welcome email",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		MaxRetries: 3,
	}
	s.Save(job1)
	fmt.Printf("[SAVE] Job created: ID=%s, Status=%s\n", job1.ID, job1.Status)

	// 2. Retrieve the job
	retrievedJob, ok := s.Get(jobID)
	if ok {
		fmt.Printf("[GET]  Retrieved: ID=%s, Payload=%q\n", retrievedJob.ID, retrievedJob.Payload)
	} else {
		fmt.Printf("[GET]  Failed to retrieve job %s\n", jobID)
	}

	// 3. Update the job status
	retrievedJob.Status = models.StatusProcessing
	retrievedJob.UpdatedAt = time.Now()
	s.Update(retrievedJob)
	fmt.Printf("[UPDATE] Status changed: %s\n", retrievedJob.Status)

	// 4. Verify the update
	updatedJob, _ := s.Get(jobID)
	fmt.Printf("[VERIFY] New Status in Store: %s\n", updatedJob.Status)

	// 5. Test missing job
	_, ok = s.Get("missing-id")
	fmt.Printf("[GET]  Missing ID check (ok=%v)\n", ok)

}
